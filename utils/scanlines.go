package utils

import (
	"bufio"
	"os"
	"strings"
)

func ScanLines() string {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return strings.Join(lines, "\n")
}
