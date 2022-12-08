package main

import (
	"github.com/hauk184/advent/files"
	"sort"
	"strings"
)

type buf struct {
	value string
	index int
}

const BUFFER_LENGTH = 14

func main() {
	input := files.ReadFile("day6/input")[0]

	buffer := make(chan buf)
	done := make(chan bool)
	go producer(input, buffer)
	go consumer(buffer, done)
	<-done
}

func dupes(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}
func consumer(buffer chan buf, done chan bool) {
	for i2 := range buffer {
		split := strings.Split(i2.value, "")
		sort.Strings(split)
		join := strings.Join(split, "")

		if !dupes(join) {
			println(i2.index)
			done <- true
		}
	}
	done <- true
}

func producer(input string, buffer chan buf) {
	for i := BUFFER_LENGTH; i < len(input); i++ {
		buffer <- buf{value: input[i-BUFFER_LENGTH : i], index: i}
	}
	close(buffer)
}
