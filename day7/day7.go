package main

import (
	"github.com/hauk184/advent/files"
	"sort"
	"strconv"
	"strings"
)

type directory struct {
	parent    *directory
	directory map[string]directory
	files     map[string]file
	size      *int
}

type file struct {
	size int
}

func main() {
	input := files.ReadFile("day7/input")

	s := 0
	root := directory{parent: nil, directory: map[string]directory{}, files: map[string]file{}, size: &s}
	parse(input, &root)
	sumFiles := sum(&root)
	println(sumFiles)
	calculateSizes(&root)
	total := 70000000
	needed := 30000000
	unused := total - *root.size
	target := needed - unused
	println(findSmallest(root, target))

}

func calculateSizes(dir *directory) {
	sum := 0
	for _, v := range dir.directory {
		calculateSizes(&v)
		sum += *v.size
	}
	for _, f := range dir.files {
		sum += f.size
	}
	dir.size = &sum
}

func findSmallest(dir directory, target int) int {
	dirchan := make(chan int)
	done := make(chan int)

	go producer(dir, dirchan)
	go consumer(dirchan, target, done)
	return <-done
}

func consumer(directories chan int, target int, done chan int) {
	var dirs []int
	for i := range directories {
		dirs = append(dirs, i)
	}

	sort.Ints(dirs)
	for _, dir := range dirs {
		if dir >= target {
			done <- dir
		}
	}

}

func producer(dir directory, dirs chan int) {
	for _, d := range dir.directory {
		dirs <- *d.size
		producer(d, dirs)
	}
	if dir.parent == nil {
		close(dirs)
	}
}

func sum(directory2 *directory) int {
	sumFiles := 0
	for _, d := range directory2.directory {
		sumFiles += sum(&d)
	}
	if *directory2.size <= 100000 {
		sumFiles += *directory2.size
	}
	return sumFiles
}

func parse(cmdList []string, current *directory) {
	if len(cmdList) == 0 {
		return
	}
	cmd := cmdList[0]
	next := cmdList[1:]
	if strings.HasPrefix(cmd, "$ cd ") {
		d := cmd[5:]
		if d == ".." {
			*(current).parent.size += *current.size
			parse(next, current.parent)
		} else if d == "/" {
			sum := 0
			newdir := directory{parent: current, directory: map[string]directory{}, files: map[string]file{}, size: &sum}
			current.directory["/"] = newdir
			parse(next, &newdir)
		} else {
			newdir := current.directory[d]
			parse(next, &newdir)
		}
	} else if strings.HasPrefix(cmd, "$ ls") {
		parse(next, current)
	} else if strings.HasPrefix(cmd, "dir") {
		dirname := cmd[4:]
		_, ok := current.directory[dirname]
		if ok {
			parse(next, current)
		} else {
			sum := 0
			(*current).directory[dirname] = directory{parent: current, directory: map[string]directory{}, files: map[string]file{}, size: &sum}
			parse(next, current)
		}
	} else {
		f := strings.Split(cmd, " ")
		atoi, _ := strconv.Atoi(f[0])
		println(len(cmdList))
		(*current).files[f[1]] = file{size: atoi}
		*current.size += atoi
		parse(next, current)
	}
}
