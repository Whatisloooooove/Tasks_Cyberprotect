package main

import (
	"fmt"
)

func main() {
	testCase1 := []string{}
	fmt.Println(Reverse(testCase1))

	testCase2 := []string{"Hello", "World"}
	fmt.Println(Reverse(testCase2))

	testCase3 := []string{"лес", "был", "тёмный"}
	fmt.Println(Reverse(testCase3))
}
