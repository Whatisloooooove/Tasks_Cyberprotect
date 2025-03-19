package sortinterface

import (
	"sort"
	"testing"
)

var testCase = []User{
	{"Лена", 25},
	{"Влад", 20},
	{"Олег", 42},
	{"Маша", 30},
	{"Никита", 13},
}

// Тест сортировки по имени
func TestSortInterfaceByName(t *testing.T) {
	sort.Sort(ByName(testCase))

	expected := []User{
		{"Влад", 20},
		{"Лена", 25},
		{"Маша", 30},
		{"Никита", 13},
		{"Олег", 42},
	}

	for i := range len(expected) {
		if expected[i] != testCase[i] {
			t.Errorf("Ожидался элемент %v, но получили %v", expected[i], testCase[i])
		}
	}
}

// Тест сортировки по возрасту
func TestSortInterfaceByAge(t *testing.T) {
	sort.Sort(ByAge(testCase))

	expected := []User{
		{"Никита", 13},
		{"Влад", 20},
		{"Лена", 25},
		{"Маша", 30},
		{"Олег", 42},
	}

	for i := range len(expected) {
		if expected[i] != testCase[i] {
			t.Errorf("Ожидался элемент %v, но получили %v", expected[i], testCase[i])
		}
	}

}
