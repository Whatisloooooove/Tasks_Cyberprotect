package stringtools

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// readAllAndRestoreStdout читает все файлы и восстанавливает Stdout
func readAllAndRestoreStdout(r *os.File, old *os.File) (string, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		return "", err
	}
	os.Stdout = old
	return buf.String(), nil
}

// Тест для проверки функции PrintHelloWorld
func TestPrintHelloWorld(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintHelloWorld()

	w.Close()
	actual, _ := readAllAndRestoreStdout(r, old)

	expected := "Hello World\n"
	if actual != expected {
		t.Errorf("Хотели вывести: %s, но получили: %s", expected, actual)
	}
}

// Тест для проверки функции SplitAndPrintWords
func TestSplitAndPrintWords(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected []string
	}{
		{
			input:    "apple,banana,orange",
			expected: []string{"apple", "banana", "orange"},
		},
		{
			input:    "dog,cat",
			expected: []string{"dog", "cat"},
		},
		{
			input:    "hello,world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{""},
		},
		{
			input:    ",",
			expected: []string{""},
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			SplitAndPrintWords(tc.input)

			w.Close()
			actual, _ := readAllAndRestoreStdout(r, old)

			lines := strings.Split(strings.TrimSpace(actual), "\n")

			if len(lines) != len(tc.expected) {
				t.Fatalf("Ожидали %d строк, получили %d", len(tc.expected), len(lines))
			}

			for i, line := range lines {
				if line != tc.expected[i] {
					t.Errorf("Строка №%d: ожидали %q, получили %q", i+1, tc.expected[i], line)
				}
			}
		})
	}
}

// Тест для проверки функции PrintCountWords
func TestPrintCountWords(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected string
	}{
		{
			input:    "Go is a programming language",
			expected: "5\n",
		},
		{
			input:    "",
			expected: "0\n",
		},
		{
			input:    "mama papa",
			expected: "2\n",
		},
		{
			input:    "a a a a a a a a a a",
			expected: "10\n",
		},
		{
			input:    ", ! %",
			expected: "3\n",
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintCountWords(tc.input)

			w.Close()
			actual, _ := readAllAndRestoreStdout(r, old)

			if actual != tc.expected {
				t.Errorf("Ожидали %q, получили %q", tc.expected, actual)
			}

		})
	}
}
