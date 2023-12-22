#include <iostream>
#include <queue>
#include <string>
#include <vector>
#include <unordered_set>

#include "match.h"

using namespace std;

class TeamMetadata{

public: 
    string name;
    int total_matches_played;

    unordered_map<time_t, int> goals_scored;
    unordered_map<time_t, int> goals_received;
    unordered_map<time_t, int> yellow_cards_received;    

    priority_queue<time_t> matches_played_date;
    unordered_set<string> matches_played_id;

    TeamMetadata(string name);
    void registerMatch(Match m, int t_order); // O(lg D)
};