#include "match.h"

Match::Match(string t1, string t2, tuple<int,int> goals, time_t date){
    this->team1 = t1;
    this->team2 = t2;
    this->goals = goals;
    this->date = date;

    this->id = date + "_" + t1 + "_" + t2;
}

bool Match::both_scored(){
    return get<0>(this->goals) > 0 && get<1>(this->goals);
}