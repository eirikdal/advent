package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content := ReadFile()

	var sums = make([]int, 2048)
	fn := 0
	for _, s := range content {
		if s != "" {
			intVal, _ := strconv.Atoi(s)
			sums[fn] = intVal + sums[fn]
		} else {
			fn++
		}
	}
	//fmt.Println(max(sums))
	sort.Ints(sums)
	fmt.Println(sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3])
}

func max(arr [2048]int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile() []string {
	file, _ := os.ReadFile("calories")
	content := strings.Split(string(file), "\n")

	return content
}

func Hello(foobar string) error {
	fmt.Println(foobar)

	return nil
}
