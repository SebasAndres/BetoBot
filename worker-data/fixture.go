package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Fixture struct {
	name    string
	year    int
	Matches LinkedList
	Teams   map[string]Team
	N_TEAMS int
	zScore  int
}

func NewFixture(name string, year int) *Fixture {
	return &Fixture{
		name:    name,
		year:    year,
		Matches: LinkedList{},
		Teams:   make(map[string]Team),
		N_TEAMS: 0,
		zScore:  0,
	}
}

func downloadFixture(rdb *redis.Client, ctx context.Context, tournament_name string, year int, url string) Fixture {

	// Leer el cuerpo del httpGet
	body := _httpGET(url)

	rdb.SAdd(ctx, PREFIX+"_URLS_DOWNLOADED", url)

	// Parsear el JSON
	var data []Match
	json.Unmarshal(body, &data)

	// Inicializar un fixture vacio
	fixture := NewFixture(tournament_name, year)

	// Agregar los partidos al fixture y a Redis
	TOURNAMENT_MATCHES_SORTED_SET := PREFIX + tournament_name + "_MATCHES"
	for _, match := range data {
		// Agregar el partido al fixture
		fixture.registerMatch(match)

		// Agregar el partido al SORTED_SET de partidos del torneo, ordenado por fecha
		unixMatchDate := convertToUnix(match.Date)
		serialized_match, _ := json.Marshal(MatchAsMap(&match))
		rdb.ZAdd(
			ctx,
			TOURNAMENT_MATCHES_SORTED_SET,
			redis.Z{
				Score:  float64(unixMatchDate),
				Member: string(serialized_match),
			},
		)

		// Agregar el partido a la estructura (arbol) de partidos

	}
	return *fixture
}

func (f *Fixture) registerMatch(match Match) {
	// Registro los equipos
	if _, exists := f.Teams[match.HomeTeam]; !exists {
		// fmt.Println("New team: " + match.HomeTeam)
		f.Teams[match.HomeTeam] = *NewTeam(match.HomeTeam)
		f.N_TEAMS++
	}
	if _, exists := f.Teams[match.AwayTeam]; !exists {
		// fmt.Println("New team: " + match.AwayTeam)
		f.Teams[match.AwayTeam] = *NewTeam(match.AwayTeam)
		f.N_TEAMS++
	}

	// Registro el partido en la lista de partidos
	f.Matches.Add(match)

	// Actualizo el zScore de la liga
	if match.HomeScore > 0 && match.AwayScore > 0 {
		f.zScore++
	}

	// Actualizo los equipos
	homeTeam := f.Teams[match.HomeTeam]
	homeTeam.NUM_PLAYED++
	homeTeam.GoalsScored += match.HomeScore
	homeTeam.GoalsReceived += match.AwayScore
	if match.HomeScore > 0 {
		homeTeam.XScore++
	}
	if match.AwayScore == 0 {
		homeTeam.QScore++
	}
	f.Teams[match.HomeTeam] = homeTeam

	awayTeam := f.Teams[match.AwayTeam]
	awayTeam.NUM_PLAYED++
	awayTeam.GoalsScored += match.AwayScore
	awayTeam.GoalsReceived += match.HomeScore
	if match.AwayScore > 0 {
		awayTeam.XScore++
	}
	if match.HomeScore == 0 {
		awayTeam.QScore++
	}
	f.Teams[match.AwayTeam] = awayTeam

	// Actualizo las estructuras de probabilidad conjunta
	// TODO
}

func (f *Fixture) estimatePZ() float64 {
	return float64(f.zScore) / float64(f.Matches.Length)
}

func (f *Fixture) getFixtureName() string {
	return f.name + "_" + strconv.Itoa(f.year)
}
