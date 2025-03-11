package main

import (
	"fmt"
	"strings"
)

func PrintHelloWorld() {
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)
}

func SplitAndPrintWords(str string) {
	parts := strings.Split(str, ",")
	for _, part := range parts {
		fmt.Println(part)
	}
}

func PrintCountWords(str string) {
	words := strings.Fields(str)
	fmt.Println(len(words))
}

func main() {
	PrintHelloWorld()
	SplitAndPrintWords("apple,banana,orange")
	PrintCountWords("Go is a programming language")

}
