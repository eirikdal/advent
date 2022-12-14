package main

import (
	"github.com/hauk184/advent/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile("day8/input")
	treeMap := make([][]int, len(input))
	treeMap_t := make([][]int, len(input))
	scenicMap := make([][]int, len(input))

	for i, _ := range treeMap {
		treeMap[i] = make([]int, len(input))
		treeMap_t[i] = make([]int, len(input))
		scenicMap[i] = make([]int, len(input))
	}

	for i, v := range input {
		str := strings.Split(v, "")
		for j, s := range str {
			num, _ := strconv.Atoi(s)

			treeMap[i][j] = num
			treeMap_t[j][i] = num
			scenicMap[i][j] = 0
		}
	}

	for i, _ := range treeMap {
		if i == 0 || i == len(input)-1 {
			continue
		}
		for j, _ := range treeMap[i] {
			if j == 0 || j == len(input)-1 {
				continue
			}
			score := calcScenicScore(treeMap[i], j)
			scenicScore := calcScenicScore(treeMap_t[j], i)
			scenicMap[i][j] = score * scenicScore
		}
	}

	maxScenicScore := 0
	for _, ints := range scenicMap {
		for _, i := range ints {
			if i > maxScenicScore {
				maxScenicScore = i
			}
		}
	}
	println(maxScenicScore)
}

func calcScenicScore(input []int, index int) int {
	scenicScoreLeft := 0
	scenicScoreRight := 0
	for i, v := range input {
		if i < index { // west, north
			if v >= input[index] {
				scenicScoreLeft = 1
			} else {
				scenicScoreLeft += 1
			}
		} else if i > index { // e
			scenicScoreRight += 1 // ast, south
			if v >= input[index] {
				break
			}
		}
	}

	return scenicScoreRight * scenicScoreLeft
}
