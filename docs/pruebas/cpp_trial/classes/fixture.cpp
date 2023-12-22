#include "fixture.h"

Fixture::Fixture(string url, string name){
    // Inicializacion de los atributos del Fixture
    this->url = url;
    this->name = name;

    // Lectura de los datos desde la url de FixtureDownload...
    vector<Match> data = getMatchesByUrl(url);
    for(Match m: data){                         // O(M * log F * log D)
        this->registerMatch(m);                 
    }
}

vector<Match> getMatchesByUrl(string url){
    // Devuelve los partidos del fixture en esa url
    return vector<Match>();
}

void Fixture::registerMatch(Match m){
    // Registra un partido en el Fixture y en cada Equipo

    if (this->matches_played.find(m.id) == this->matches_played.end()){

        // Agrego el partido en el Fixture
        this->matches_played.insert(m.id);
        if (this->set_dates.find(m.date) == this->set_dates.end()){
            this->dates.push(m.date);   // O(lg F), F: #fechas del fixture
            this->set_dates.insert(m.date);
        }

        // Agrego los datos a las estructuras del Fixture (both_scored, yellow_cards, etc)
        if (bothScoredByMatch.find(m.date) == bothScoredByMatch.end()){
            vector<bool> v = {};
            v.push_back(m.both_scored());
            this->bothScoredByMatch.insert({m.date, v});
        }
        else {
            bothScoredByMatch[m.date].push_back(m.both_scored());
        }

        // Agregamos los datos a las estructuras internas de los equipos
        // [EQ1] Si el equipo1 no esta en teams Metadata
        if (this->teams.find(m.team1) == this->teams.end()){    
            TeamMetadata team_metadata = TeamMetadata(m.team1);
            team_metadata.registerMatch(m, 0);                          // O(lg D)
            this->teams.insert({m.team1, team_metadata});
        }
        else {
            this->teams[m.team1].registerMatch(m, 0);
        }
        // [EQ2] Si el equipo2 no esta en teams Metadata
        if (this->teams.find(m.team2) == this->teams.end()){
            TeamMetadata team_metadata = TeamMetadata(m.team2);
            team_metadata.registerMatch(m,1);                           // O(lg D)
            this->teams.insert({m.team2, team_metadata});
        }
        else {
            this->teams[m.team2].registerMatch(m, 1);
        }
    }
    else {
        // El partido ya esta registrado
    }
}

void Fixture::uploadToRedis(){
    // Subir las estructuras de Equipo
    
    // Subir las estructuras de la Liga
    
}