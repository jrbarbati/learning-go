package main

type Score struct {
	TeamName string
	score    int
}

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Name  string
	Teams []Team
	Wins  map[string]int
}

func NewScore(name string, score int) Score {
	return Score{
		TeamName: name,
		score:    score,
	}
}

func NewLeague() League {
	return League{
		Teams: make([]Team, 0),
		Wins:  make(map[string]int),
	}
}

func NewTeam(name string, playerNames []string) Team {
	return Team{
		Name:        name,
		PlayerNames: playerNames,
	}
}

func main() {

}
