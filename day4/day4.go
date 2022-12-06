package main

import (
	"github.com/hauk184/advent/files"
	"strconv"
	"strings"
)

type sections struct {
	first  rng
	second rng
}

type rng struct {
	lower  int
	higher int
}

func main() {
	contents := files.ReadFile("day4/input")

	part1(contents)
	part2(contents)
}

func part2(contents []string) {
	count := 0
	for _, v := range contents {
		section := toSection(v)

		if overlaps(section.first, section.second) ||
			overlaps(section.second, section.first) {
			count += 1
		}
	}

	println(count)
}

func part1(contents []string) {
	count := 0
	for _, v := range contents {
		section := toSection(v)

		if contains(section.first, section.second) ||
			contains(section.second, section.first) {
			count += 1
		}
	}

	println(count)
}
func overlaps(first rng, other rng) bool {
	if first.lower >= other.lower && first.lower <= other.higher {
		return true
	} else if first.higher <= other.higher && first.higher >= other.lower {
		return true
	}
	return false
}
func contains(first rng, other rng) bool {
	if first.lower >= other.lower && first.higher <= other.higher {
		return true
	}
	return false
}

func toSection(str string) sections {
	split := strings.Split(str, ",")

	return sections{first: toRange(split[0]), second: toRange(split[1])}
}

func toRange(str string) rng {
	r := strings.Split(str, "-")

	lower, _ := strconv.Atoi(r[0])
	higher, _ := strconv.Atoi(r[1])

	return rng{lower: lower, higher: higher}
}
