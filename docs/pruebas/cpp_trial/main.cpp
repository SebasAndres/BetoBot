#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
// #include <curl/curl.h>

#include "classes/match.h"
#include "classes/fixture.h"

using namespace std;

struct webFixture{
    string url;
    string name;
}

int main() { 
    vector<webFixture> web_fixtures = {
        (webFixture){"messi.com", "Argentina"}            
    }; 
    
    for (webFixture league: web_fixtures){                    // O(N * f)
        
        // Por cada liga      
        cout << '-----------' << endl;              
        cout << league.name << endl;

        // Leo los datos de la FixtureDownload : O(1)
        Fixture fixtureData = Fixture(league.url,  
                                      league.name);
        
        // Actualizo el Redis                  : O(?)
        updateRedisData(fixtureData);                       
    }    

    return 0;
}

void updateRedisData(Fixture fixtureData){
    // 
}