package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {

	// Conexion a Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADRESS,
		Password: REDIS_PASSWORD,
		DB:       0, // default DB
	})
	ctx := context.Background()

	// Levantar las ligas de un archivo externo
	// Para que no haya que recompilar el programa cada vez que se quiera agregar una liga
	// Cada liga tiene un nombre y una URL y se agrega en el SET "FB_LEAGUES" de Redis
	var tournaments map[string]map[int]string
	tournaments = map[string]map[int]string{
		"CHAMPIONS": {
			2022: "https://fixturedownload.com/feed/json/champions-league-2022",
			2021: "https://fixturedownload.com/feed/json/champions-league-2021",
			//2020: "https://fixturedownload.com/feed/json/champions-league-2020",
		},
	}

	// Por cada tournaments
	for tournament, urls := range tournaments {
		log.Printf("Processing %s", tournament)

		// Agrego la liga al SET de ligas procesadas en Redis
		rdb.SAdd(ctx, PREFIX+"TOURNAMENTS", tournament)

		cup_total_matches := 0
		var cup_Z []int

		// Por cada fixture del torneo
		for year, url := range urls {

			if rdb.SIsMember(ctx, PREFIX+"_URLS_DOWNLOADED", url).Val() {
				fmt.Println("Fixture " + tournament + "_" + strconv.Itoa(year) + " already downloaded")
				continue
			}

			fmt.Println("Downloading " + strconv.Itoa(year))
			fmt.Println("From: " + url)

			// Descargar el fixture
			fixture := downloadFixture(rdb, ctx, tournament, year, url)

			// Actualizar datos procesados en Redis tras la lectura de todos los partidos
			fixture_pz := fixture.estimatePZ()
			cup_total_matches += fixture.Matches.Length
			cup_Z = append(cup_Z, int(fixture_pz*float64(fixture.Matches.Length)))
			rdb.HSet(ctx, PREFIX+tournament, "pZ_"+strconv.Itoa(year), fixture_pz).Err()

			// Por cada equipo en el fixture, actualizo las estructuras
			for _, team := range fixture.Teams {

				// Agrego el equipo al cjto de equipos en Redis
				rdb.SAdd(ctx, PREFIX+"TEAMS", team.Name)

				team_varname := PREFIX + "TEAM_" + strings.ToUpper(team.Name)
				// Update Tournament Team Data
				e := rdb.HIncrBy(ctx, team_varname, tournament+"_NUM_PLAYED", int64(team.NUM_PLAYED)).Err()
				if e != nil {
					log.Println(e)
				}
				rdb.HIncrBy(ctx, team_varname, tournament+"_XScore", int64(team.XScore)).Err()
				rdb.HIncrBy(ctx, team_varname, tournament+"_QScore", int64(team.QScore)).Err()
				rdb.HIncrBy(ctx, team_varname, tournament+"_GoalsScored", int64(team.GoalsScored)).Err()
				rdb.HIncrBy(ctx, team_varname, tournament+"_GoalsReceived", int64(team.GoalsReceived)).Err()

				// Update Global Team Data
				rdb.HIncrBy(ctx, team_varname, "NUM_PLAYED", int64(team.NUM_PLAYED)).Err()
				rdb.HIncrBy(ctx, team_varname, "XScore", int64(team.XScore)).Err()
				rdb.HIncrBy(ctx, team_varname, "QScore", int64(team.QScore)).Err()
				rdb.HIncrBy(ctx, team_varname, "GoalsScored", int64(team.GoalsScored)).Err()
				rdb.HIncrBy(ctx, team_varname, "GoalsReceived", int64(team.GoalsReceived)).Err()
			}

		}

		// Agrego los datos generales del torneo
		var cup_pZ float64 = 0
		for _, z := range cup_Z {
			cup_pZ += float64(z)
		}
		cup_pZ /= float64(cup_total_matches)
		rdb.HSet(ctx, PREFIX+tournament, "pZ", cup_pZ).Err()
	}
}

func convertToUnix(date string) int64 {
	layout := "2006-01-02 15:04:05Z"
	t, _ := time.Parse(layout, date)
	return t.Unix()
}
