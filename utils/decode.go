package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Decode(input string) (string, error) {

	if !BalanceBracket(input) {
		return "", errors.New("Error: Square brackets are unbalanced")
	}

	// {	re := regexp.MustCompile(`\[(\d+) ([^\]]+)\]`)

	re := regexp.MustCompile(`\[(.*?)\]`)
	var decodeErr error

	output := re.ReplaceAllStringFunc(input, func(match string) string {
		content := match[1 : len(match)-1]

		spaceIndex := strings.Index(content, " ")
		if spaceIndex == -1 {
			decodeErr = errors.New("Error: The arguments are not separated by a space")
			return ""
		}
		parts := strings.SplitN(content, " ", 2)

		if len(parts) != 2 {
			decodeErr = errors.New("Error: There is not enough arguments")
			return ""
		}
		firstArg := content[:spaceIndex]
		secondArg := content[spaceIndex+1:]

		if secondArg == "" {
			decodeErr = errors.New("Error: There is no second argument")
			return ""
		}

		count, err := strconv.Atoi(firstArg)
		if err != nil {
			decodeErr = errors.New("Error: The first argument is not a number")
			return ""
		}

		return strings.Repeat(secondArg, count)
	})

	if decodeErr != nil {
		return "", decodeErr
	}

	return output, nil
}
