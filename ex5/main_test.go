package ex5_test

import (
	"strings"
	"testing"
)

const sentence = "Education is what remains after one has forgotten what one has learned in school."

var result string

func BenchmarkToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = strings.ToUpper(sentence)
	}
}
