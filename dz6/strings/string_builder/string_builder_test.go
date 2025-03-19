package stringbuilder

import (
	"strings"
	"testing"
)

// Тест работы структуры StringBuilder
func TestStringBuilder(t *testing.T) {
	sb := &StringBuilder{}
	expected := "Hello, World!"
	sb.WriteString("Hello, ")
	sb.WriteString("World!")

	if sb.String() != expected {
		t.Errorf("Ожидалось: %q, получилось: %q", expected, sb.String())
	}
}

// Бенчмарк для реализации StringBuilder
func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sb := &StringBuilder{}
		sb.WriteString("Hello, ")
		sb.WriteString("World!")
		_ = sb.String()
	}
}

// Бенчмарк для стандартного strings.Builder
func BenchmarkStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.WriteString("Hello, ")
		sb.WriteString("World!")
		_ = sb.String()
	}
}
