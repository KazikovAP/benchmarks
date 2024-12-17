package benchmarks_test

import (
	"strings"
	"testing"

	"github.com/KazikovAP/benchmarks/benchmarks"
)

func generateStringWithSubstring(sizeInBytes int) string {
	unit := strings.Repeat("a", sizeInBytes)
	return strings.Repeat(unit+"b", 1000)
}

func BenchmarkCount(b *testing.B) {
	cases := []struct {
		name     string
		substr   string
		size     int
		function func(string, string) int
	}{
		{"Strings_6", "abcdef", 6, benchmarks.CountWithStrings},
		{"Regex_6", "abcdef", 6, benchmarks.CountWithRegex},
		{"Manual_6", "abcdef", 6, benchmarks.CountManually},
		{"Strings_16", "abcdefghijklmnop", 16, benchmarks.CountWithStrings},
		{"Regex_16", "abcdefghijklmnop", 16, benchmarks.CountWithRegex},
		{"Manual_16", "abcdefghijklmnop", 16, benchmarks.CountManually},
		{"Strings_26", "abcdefghijklmnopqrstuvwxyz", 26, benchmarks.CountWithStrings},
		{"Regex_26", "abcdefghijklmnopqrstuvwxyz", 26, benchmarks.CountWithRegex},
		{"Manual_26", "abcdefghijklmnopqrstuvwxyz", 26, benchmarks.CountManually},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			str := generateStringWithSubstring(tc.size)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tc.function(str, tc.substr)
			}
		})
	}
}
