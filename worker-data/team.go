package main

type Team struct {
	Name string
	// En todos los torneos
	NUM_PLAYED    int
	XScore        int // # of matches with at least one goal scored
	QScore        int // # of matches with no goals scored
	GoalsScored   int
	GoalsReceived int
}

func NewTeam(name string) *Team {
	return &Team{
		Name:          name,
		NUM_PLAYED:    0,
		XScore:        0,
		QScore:        0,
		GoalsScored:   0,
		GoalsReceived: 0,
	}
}

func FixtureTeamAsMap(t *Team, tournament string) map[string]interface{} {
	return map[string]interface{}{
		tournament + "_NUM_PLAYED":    t.NUM_PLAYED,
		tournament + "_XScore":        t.XScore,
		tournament + "_QScore":        t.QScore,
		tournament + "_GoalsScored":   t.GoalsScored,
		tournament + "_GoalsReceived": t.GoalsReceived,
	}
}
