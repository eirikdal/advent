package files

import (
	"os"
	"strings"
)

func ReadFile() []string {
	file, _ := os.ReadFile("calories")
	content := strings.Split(string(file), "\n")

	return content
}
