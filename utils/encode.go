package utils

import (
	"fmt"
	"strings"
)

func Encode(input string, maxPattern int) (string, error) {
	if input == "" {
		return "", nil
	}

	var result strings.Builder

	for i := 0; i < len(input); {
		bestPattern := ""
		bestCount := 1
		bestLength := 0

		remaining := input[i:]

		for size := 1; size <= maxPattern && size <= len(remaining)/2; size++ {
			pattern := remaining[:size]
			count := 1

			for (count+1)*size <= len(remaining) &&
				remaining[count*size:(count+1)*size] == pattern {
				count++
			}

			if count > 1 {
				originalLen := count * size

				// single-char runs: encode only if long enough
				if size == 1 && count < 3 {
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
			encodedStr := fmt.Sprintf("[%d %s]", bestCount, bestPattern)
			result.WriteString(encodedStr)
			i += bestCount * len(bestPattern)
		} else {
			result.WriteByte(input[i])
			i++
		}
	}

	return result.String(), nil
}
