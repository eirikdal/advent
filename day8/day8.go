package main

import (
	"github.com/hauk184/advent/files"
	"strconv"
	"strings"
)

type tree struct {
	index   int
	visible bool
}

func main() {
	input := files.ReadFile("day8/input")
	treeMap := make([][]int, len(input))
	treeMap_t := make([][]int, len(input))
	visibleMap := make([][]bool, len(input))

	for i, _ := range treeMap {
		treeMap[i] = make([]int, len(input))
		treeMap_t[i] = make([]int, len(input))
		visibleMap[i] = make([]bool, len(input))
	}

	for i, v := range input {
		str := strings.Split(v, "")
		for j, s := range str {
			num, _ := strconv.Atoi(s)

			treeMap[i][j] = num
			treeMap_t[j][i] = num
			visibleMap[i][j] = true
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
			visibleMap[i][j] = isVisible(treeMap[i], j) || isVisible(treeMap_t[j], i)
		}
	}

	sumVisible := 0
	for _, bools := range visibleMap {
		for _, b := range bools {
			if b {
				sumVisible += 1
				print(1)
			} else {
				print(0)
			}
		}
		println()
	}
	print(sumVisible)
}

func isVisible(input []int, index int) bool {
	visibleRight := true
	visibleLeft := true
	for i, v := range input {
		//if i != index {
		if i < index { // west, north
			if v >= input[index] {
				visibleLeft = false
			}
		} else if i > index { // east, south
			if v >= input[index] {
				visibleRight = false
			}
		}
		//}
	}

	return visibleRight || visibleLeft
}
