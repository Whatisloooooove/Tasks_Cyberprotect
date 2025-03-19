package stringtools

import (
	"fmt"
	"strings"
)

// PrintHelloWorld выводит Hello World
func PrintHelloWorld() {
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)
}

// SplitAndPrintWords нарезает строку по запятой
// и выводит каждое слово с новой строки
func SplitAndPrintWords(str string) {
	parts := strings.Split(str, ",")
	for _, part := range parts {
		fmt.Println(part)
	}
}

// PrintCountWords выводит сколько слов находится в строке
func PrintCountWords(str string) {
	words := strings.Fields(str)
	fmt.Println(len(words))
}
