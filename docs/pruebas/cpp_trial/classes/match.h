#include <iostream>
#include <string>
#include <tuple>

using namespace std;

class Match {
    public: 
        string team1;
        string team2;
        string id;
        time_t date;

        tuple<int, int> goals;
        tuple<int, int> yellow_cards;

        Match(string t1, string t2, tuple<int,int> goals, time_t date);
        bool both_scored();
};

