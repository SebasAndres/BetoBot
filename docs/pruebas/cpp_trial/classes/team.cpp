#include "team.h"

TeamMetadata::TeamMetadata(string name){
    this->name = name;
}

void TeamMetadata::registerMatch(Match m, int team_order){
    // Registramos un partido en la estructura del equipo en la posicion team_order del partido

    if (this->matches_played_id.find(m.id) == this->matches_played_id.end()){
        // Actualizamos cantidad y partidos jugados
        this->total_matches_played++;
        this->matches_played_date.push(m.date);
        this->matches_played_id.insert(m.id); // O(lg D), D: # fechas registradas

        if (team_order == 0){
            // Registramos goles marcados, recibidos, amarillas, etc..
            this->goals_scored.push(get<0>(m.goals));
            this->goals_received.push(get<1>(m.goals));
        }
        else {
            // Registramos goles marcados, recibidos, amarillas, etc..
            this->goals_scored.insert(get<1>(m.goals));
            this->goals_received.push(get<0>(m.goals));
        }    
    }
    else {
        // Not a valid match (ya fue registrado)
    }

}