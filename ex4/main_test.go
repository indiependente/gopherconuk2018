package src_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const x = "hello"

var result string

func withPlus() string {
	return x + x + x + x + x + x + x + x + x + x
}

func BenchmarkWithPlus(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = withPlus()
	}
	result = s
}

func withSprintf() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", x, x, x, x, x, x, x, x, x, x)
}

func BenchmarkWithSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = withSprintf()
	}
	result = s
}

func withBuffer() string {
	bb := &bytes.Buffer{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func BenchmarkWithBuffer(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = withBuffer()
	}
	result = s
}

func withStringBuilder() string {
	bb := &strings.Builder{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func BenchmarkWithStringBuilder(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = withStringBuilder()
	}
	result = s
}
