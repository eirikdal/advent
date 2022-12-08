package main

import (
	"github.com/hauk184/advent/files"
	"github.com/hauk184/advent/stack"
	"strconv"
	"strings"
)

type crate struct {
	value string
}
type crates stack.Stack[crate]

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

func doMove(crates []stack.Stack[crate], move move) {
	a := crates[move.from].Pop(move.number)
	crates[move.to].Multipush(a)
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

func toCrates(input []string) []stack.Stack[crate] {
	crates := make([]stack.Stack[crate], 10)
	for i, _ := range crates {
		crates[i] = make(stack.Stack[crate], 0)
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
