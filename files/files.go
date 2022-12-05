package files

import (
	"os"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	content := strings.Split(string(file), "\n")

	return content, err
}
