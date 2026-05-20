package utils

import (
	"errors"
	"strings"
)

func Multiline(text string, encode bool) (string, error) {
	if text == "" {
		return "", errors.New("Error: Input cannot be empty.")
	}
	lines := strings.Split(text, "\n")

	var processedLines []string

	for _, line := range lines {
		var rowOutput string
		var err error
		if encode {
			rowOutput, err = Encode(line)
		} else {
			rowOutput, err = Decode(line)
		}

		if err != nil {
			return "", err
		}
		processedLines = append(processedLines, rowOutput)
	}
	output := strings.Join(processedLines, "\n")

	return output, nil
}
