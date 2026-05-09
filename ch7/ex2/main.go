package main

import "fmt"

type Score struct {
	TeamName string
	score    int
}

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
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
func (l *League) MatchResult(score1, score2 Score) {
	if score1.score > score2.score {
		l.Wins[score1.TeamName]++
		return
	}

	if score1.score < score2.score {
		l.Wins[score2.TeamName]++
	}
}

func (l *League) Ranking() []string {
	rankings, wins := l.splitRankingAndWins()

	for i := range len(wins) - 1 {
		for j := i + 1; j < len(wins); j++ {
			if wins[i] <= wins[j] {
				continue
			}

			tempWins := wins[i]
			wins[i] = wins[j]
			wins[j] = tempWins

			tempRankings := rankings[i]
			rankings[i] = rankings[j]
			rankings[j] = tempRankings
		}
	}

	return rankings
}

func (l *League) splitRankingAndWins() (rankings []string, wins []int) {
	rankings = make([]string, len(l.Teams))
	wins = make([]int, len(l.Teams))

	for i, team := range l.Teams {
		if ws, ok := l.Wins[team.Name]; ok {
			rankings[i] = team.Name
			wins[i] = ws
		} else {
			rankings[i] = team.Name
			wins[i] = 0
		}
	}

	return rankings, wins
}

func main() {
	l := NewLeague()
	l.Teams = []Team{
		NewTeam("Oh My!", []string{}),
		NewTeam("Lions", []string{}),
		NewTeam("Bears", []string{}),
		NewTeam("Tigers", []string{}),
	}

	l.MatchResult(
		NewScore("Lions", 12),
		NewScore("Tigers", 31),
	)

	l.MatchResult(
		NewScore("Lions", 12),
		NewScore("Bears", 31),
	)

	l.MatchResult(
		NewScore("Lions", 12),
		NewScore("Oh My!", 31),
	)

	l.MatchResult(
		NewScore("Tigers", 12),
		NewScore("Bears", 31),
	)

	l.MatchResult(
		NewScore("Tigers", 12),
		NewScore("Oh My!", 31),
	)

	l.MatchResult(
		NewScore("Bears", 12),
		NewScore("Oh My!", 31),
	)

	fmt.Println(l.Ranking())
}
