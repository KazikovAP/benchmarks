package benchmarks

import (
	"regexp"
	"strings"
)

func CountWithStrings(s, substr string) int {
	return strings.Count(s, substr)
}

func CountWithRegex(s, substr string) int {
	re := regexp.MustCompile(substr)
	return len(re.FindAllStringIndex(s, -1))
}

func CountManually(s, substr string) int {
	count, n := 0, len(substr)
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == substr {
			count++
		}
	}

	return count
}
