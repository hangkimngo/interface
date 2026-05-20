package utils

import (
	"strconv"
	"strings"
)

func Encode(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	var result strings.Builder

	for i := 0; i < len(input); {
		bestPattern := ""
		bestCount := 1
		bestLength := 0

		remaining := input[i:]

		for size := 1; size <= 2 && size <= len(remaining)/2; size++ {
			pattern := remaining[:size]

			if strings.ContainsAny(pattern, "[]") {
				continue
			}

			count := 1

			for (count+1)*size <= len(remaining) &&
				remaining[count*size:(count+1)*size] == pattern {
				count++
			}

			if count > 1 {
				originalLen := count * size

				// single-char runs: encode only if long enough
				if size == 1 && count < 4 {
					continue
				}

				// two-char patterns: encode from 2+
				if originalLen > bestLength {
					bestPattern = pattern
					bestCount = count
					bestLength = originalLen
				}
			}
		}

		if bestCount > 1 {
			result.WriteString("[")
			result.WriteString(strconv.Itoa(bestCount))
			result.WriteString(" ")
			result.WriteString(bestPattern)
			result.WriteString("]")

			i += bestCount * len(bestPattern)
		} else {
			result.WriteByte(input[i])
			i++
		}
	}

	return result.String(), nil
}
