#include <iostream>
#include <vector>
#include <unordered_map>
#include <list>

#include "match.h"
#include "team.h"

using namespace std;

class Fixture{
    private:
        string url;
        string name;

        unordered_map<string, TeamMetadata> teams;     // byTeamId, team data

        unordered_map<time_t, vector<bool>> bothScoredByMatch;              // byDateTime
        unordered_map<time_t, vector<tuple<int,int>>> yellowCardsByMatch;   // byDateTime                
 
        unordered_set<string> matches_played;                               // byMatchId
        unordered_set<time_t> set_dates;
        priority_queue<time_t, vector<int>, greater<int>> dates;

    public:
        Fixture(string url, string name); // O(M * log F * log D)
        void registerMatch(Match m); // O(log F + log D)
        void uploadToRedis();
};