package sprintf

import (
	"testing"
)

// Тест метода String
func TestPersonString(t *testing.T) {
	tests := []struct {
		input    Person
		expected string
	}{
		{Person{Name: "Alice", Age: 30}, "Name: Alice, Age: 30"},
		{Person{Name: "Bob", Age: 25}, "Name: Bob, Age: 25"},
		{Person{Name: "John", Age: 40}, "Name: John, Age: 40"},
		{Person{Name: "", Age: 0}, "Name: , Age: 0"},
	}

	for _, test := range tests {
		t.Run(test.input.Name, func(t *testing.T) {
			actual := test.input.String()

			if actual != test.expected {
				t.Errorf("Ожидали %q, получили %q", test.expected, actual)
			}
		})
	}
}
