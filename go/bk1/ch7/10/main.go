package main

import (
	"io"
	"sort"
	"os"
)

type Team struct{
	teamName string
	playerNames []string
}

type League struct{
	Name string
	Teams map[string]Team
	Wins map[string]int
}

func (l *League) MatchResult(firstTeam string, firstTeamScore int, secondTeam string, secondTeamScore int) {
	if _,ok := l.Teams[firstTeam]; !ok {
		return
	}
	if _,ok := l.Teams[secondTeam]; !ok {
		return
	}
	if firstTeamScore == secondTeamScore {
		return
	}
	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeam]++
	}
	if firstTeamScore < secondTeamScore {
		l.Wins[secondTeam]++
	}
}

func (l *League) Ranking() []string{
	winners := make([]string,0, len(l.Wins))
	for key := range l.Wins {
		winners = append(winners, key)
	}
	sort.Slice(winners, func(i,j int) bool{
		return l.Wins[winners[i]] < l.Wins[winners[j]]
	})
	return winners
}

type Ranker interface{
	Ranking()[]string
}

func RankerPrinter(r Ranker, iW io.Writer){
	rankedTeams := r.Ranking()
	for _,value := range rankedTeams {
		io.WriteString(iW, value + "\n")
	}
}

func main() {
	ug := Team{
		teamName: "Uganda",
		playerNames: []string{"1","2", "3"},
	}
	kn := Team{
		teamName: "Kenya",
		playerNames: []string{"4","5", "6"},
	}
	tz := Team{
		teamName: "Tanzania",
		playerNames: []string{"7","8", "9"},
	}

	l := &(League{
		Teams: map[string]Team{
			"Uganda": ug,
			"Kenya": kn,
			"Tanzania": tz,
		},
		Wins: make(map[string]int),
	})

	l.MatchResult("Uganda", 1, "Kenya", 2)
	l.MatchResult("Tanzania", 3, "Kenya", 2)
	l.MatchResult("Uganda", 4, "Tanzania", 2)
	l.MatchResult("Uganda", 2, "Kenya", 2)
	l.MatchResult("Tanzania", 3, "Kenya", 2)
	l.MatchResult("Uganda", 4, "Tanzania", 2)
	l.Ranking()
	RankerPrinter(l, os.Stdout)
}

