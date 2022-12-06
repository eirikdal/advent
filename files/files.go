package files

import (
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	file, err := os.ReadFile(filename)
	content := strings.Split(string(file), "\n")
	if err != nil {
		panic(err)
	}
	return content
}
