package processvalue

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
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

// Тест для проверки функции ProcessValue
func TestProcessValue(t *testing.T) {
	for _, tc := range []struct {
		input    interface{}
		expected string
	}{
		{
			input:    "привет",
			expected: "ПРИВЕТ\n",
		},
		{
			input:    uint64(100),
			expected: "200\n",
		},
		{
			input:    uint32(100),
			expected: "200\n",
		},
		{
			input:    uint16(100),
			expected: "200\n",
		},
		{
			input:    uint8(100),
			expected: "200\n",
		},
		{
			input:    uint(100),
			expected: "200\n",
		},
		{
			input:    int(100),
			expected: "200\n",
		},
		{
			input:    int8(2),
			expected: "4\n",
		},
		{
			input:    int16(100),
			expected: "200\n",
		},
		{
			input:    int32(100),
			expected: "200\n",
		},
		{
			input:    int64(100),
			expected: "200\n",
		},
		{
			input:    float32(100),
			expected: "200\n",
		},
		{
			input:    float64(100),
			expected: "200\n",
		},
		{
			input:    complex64(1.0 + 3i),
			expected: "(2+6i)\n",
		},
		{
			input:    complex128(1.0 + 3i),
			expected: "(2+6i)\n",
		},
		{
			input:    true,
			expected: "Неизвестный тип\n",
		},
	} {
		t.Run(fmt.Sprintf("%v", reflect.TypeOf(tc.input)), func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			ProcessValue(tc.input)

			w.Close()
			actual, _ := readAllAndRestoreStdout(r, old)

			if actual != tc.expected {
				t.Errorf("Ожидался вывод: %s, получили: %s", tc.expected, actual)
			}
		})
	}
}
