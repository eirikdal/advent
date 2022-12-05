package files

import (
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	file, _ := os.ReadFile(filename)
	content := strings.Split(string(file), "\n")

	return content
}
