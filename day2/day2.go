package main

import (
	"github.com/hauk184/advent/files"
)

// 1 for rock
// 2 for paper
// 3 for sciccors

// 0 for loss
// 3 for draw
// 6 for win

func main() {
	content := files.ReadFile("input")

	for _, line := range content {
		println(line)
	}
}

func OpponentIsRockMap(play string) int {
	rpsR := make(map[string]int)

	rpsR["scissors"] = 0
	rpsR["paper"] = 6
	rpsR["rock"] = 3

	return rpsR[play]
}

func PointsByPlay(youPlayed string) int {
	play := make(map[string]int)

	play["scissors"] = 3
	play["paper"] = 2
	play["rock"] = 1

	return play[youPlayed]
}
