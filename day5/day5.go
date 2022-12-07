package main

import (
	"github.com/hauk184/advent/files"
	"strconv"
	"strings"
)

type crate struct {
	value string
}
type crates []crate

type move struct {
	number int
	from   int
	to     int
}

func main() {
	input := files.ReadFile("day5/input")

	crates := toCrates(input[:10])
	moves := toMoves(input[10:])

	for _, m := range moves {
		doMove(crates, m)
	}
	for _, c := range crates {
		if len(c) == 0 {
			continue
		}
		print(c[0].value)
	}
}

func doMove(crates []crates, move move) {
	a := crates[move.from].pop(move.number)
	crates[move.to].push(a)
}

func (crates *crates) push(c crates) {
	arr := make([]crate, 0)
	arr = append(arr, c...)
	*crates = append(arr, *crates...)
}

func (crates *crates) pop(number int) crates {
	el := (*crates)[0:number]
	*crates = (*crates)[number:]
	return el
}

func toMoves(input []string) []move {
	moves := make([]move, len(input))
	for i, s := range input {
		moves[i] = toMove(s)
	}
	return moves
}

func toMove(s string) move {
	arr := strings.Split(s, " ")

	atoi, _ := strconv.Atoi(arr[1])
	from, _ := strconv.Atoi(arr[3])
	to, _ := strconv.Atoi(arr[5])

	return move{number: atoi, from: from, to: to}
}

func toCrates(input []string) []crates {
	crates := make([]crates, 10)
	for i, _ := range crates {
		crates[i] = make([]crate, 0)
	}

	for l, s := range input {
		if l >= 8 {
			break
		}

		for i := 4; i < len(s)+2; i += 4 {
			c := toCrate(s[i-4 : i-1])

			if c.value != "" {
				crates[i/4] = append(crates[i/4], c)
			}
		}
	}

	return crates
}

func toCrate(input string) crate {
	return crate{value: strings.Trim(input, " ")}
}
