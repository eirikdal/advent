package main

import (
	"fmt"
	"github.com/hauk184/advent/files"
	"strings"
)

// 1 for rock
// 2 for paper
// 3 for sciccors

// 0 for loss
// 3 for draw
// 6 for win

func main() {
	content := files.ReadFile("day2/input")
	scoreMap := GetScoreMap()
	playMap := GetPointsByPlay()
	playBook := GetPlaybook()

	score := 0
	for _, line := range content {
		play := strings.Fields(line)
		opp := play[0]
		strategy := play[1]
		toPlay := playBook[strategy][opp]
		score += scoreMap[opp][toPlay] + playMap[toPlay]
	}
	fmt.Println(score)
}

func GetPlaybook() map[string]map[string]string {
	playbook := make(map[string]map[string]string)
	playbook["X"] = make(map[string]string)
	playbook["Y"] = make(map[string]string)
	playbook["Z"] = make(map[string]string)

	playbook["X"]["A"] = "Z"
	playbook["X"]["B"] = "X"
	playbook["X"]["C"] = "Y"
	playbook["Y"]["A"] = "X"
	playbook["Y"]["B"] = "Y"
	playbook["Y"]["C"] = "Z"
	playbook["Z"]["A"] = "Y"
	playbook["Z"]["B"] = "Z"
	playbook["Z"]["C"] = "X"
	return playbook
}

// A Y
// B X
// C Z
func GetScoreMap() map[string]map[string]int {
	score := make(map[string]map[string]int)

	score["A"] = make(map[string]int)
	score["B"] = make(map[string]int)
	score["C"] = make(map[string]int)

	// rock = A || X, scissors = B || Y, paper = C || Z
	score["A"]["X"] = 3
	score["A"]["Y"] = 6
	score["A"]["Z"] = 0
	score["B"]["X"] = 0
	score["B"]["Y"] = 3
	score["B"]["Z"] = 6
	score["C"]["X"] = 6
	score["C"]["Y"] = 0
	score["C"]["Z"] = 3

	return score
}

func GetPointsByPlay() map[string]int {
	play := make(map[string]int)

	play["X"] = 1
	play["Y"] = 2
	play["Z"] = 3

	return play
}
