package main

type Match struct {
	MatchNumber int    `json:"MatchNumber"`
	RoundNumber int    `json:"RoundNumber"`
	HomeTeam    string `json:"HomeTeam"`
	AwayTeam    string `json:"AwayTeam"`
	HomeScore   int    `json:"HomeTeamScore"`
	AwayScore   int    `json:"AwayTeamScore"`
	Date        string `json:"DateUtc"`
	Tournament  string
}

func setTournament(m *Match, tournament string) {
	m.Tournament = tournament
}

func MatchAsMap(m *Match) map[string]interface{} {
	return map[string]interface{}{
		"MatchNumber": m.MatchNumber,
		"RoundNumber": m.RoundNumber,
		"HomeTeam":    m.HomeTeam,
		"AwayTeam":    m.AwayTeam,
		"HomeScore":   m.HomeScore,
		"AwayScore":   m.AwayScore,
		"Date":        m.Date,
	}
}
