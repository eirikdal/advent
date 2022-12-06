package main

import (
	"github.com/hauk184/advent/files"
)

func main() {
	contents := files.ReadFile("input")

	//part1(contents)
	part2(contents)
}

func part2(contents []string) {
	sum := 0

	for i := 0; i < len(contents); i += 3 {
		first := GetValueMap(contents[i])
		snd := GetValueMap(contents[i+1])
		thrd := GetValueMap(contents[i+2])

		hits1 := make(map[int]int, 53)
		for i := 0; i < len(first); i++ {
			hits1[int(first[i])] += 1
		}
		hits2 := make(map[int]int, 53)
		for i := 0; i < len(snd); i++ {
			hits2[int(snd[i])] += 1
		}

		hits3 := make(map[int]int, 53)
		for i := 0; i < len(thrd); i++ {
			hits3[int(thrd[i])] += 1
		}

		for i := 0; i < 53; i++ {
			if hits1[i] > 0 && hits2[i] > 0 && hits3[i] > 0 {
				sum += i
				break
			}
		}
	}

	println(sum)
}

func part1(contents []string) {
	sum := 0
	for _, content := range contents {
		half := len(content) / 2

		left := GetValueMap(content[:half])
		right := GetValueMap(content[half:])

		hits := make(map[int]int, 53)
		for i := 0; i < len(left); i++ {
			hits[int(left[i])] += 1
		}
		for i := 0; i < len(right); i++ {
			hasHits := hits[int(right[i])] > 0
			if hasHits {
				sum += int(right[i])
				break
			}
		}
	}
	println(sum)
}

// uppercase = -64
// lowercase = -96
func GetValueMap(half string) []int32 {
	valueMap := make([]int32, len(half))
	for i, i3 := range half {
		if i3 >= 65 && i3 <= 90 {
			valueMap[i] = i3 - (64 - 26)
		} else if i3 >= 97 && i3 <= 122 {
			valueMap[i] = i3 - 96
		}
	}
	return valueMap
}
