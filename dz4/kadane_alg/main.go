package main

import (
	"fmt"
)

func main() {
	testCase1 := []int{1, -2, 5, -4, 30, -24}
	fmt.Println(FindSubSumKadane(testCase1))

	testCase2 := []int{-1, -2, -3, -4, -5}
	fmt.Println(FindSubSumKadane(testCase2))

	testCase3 := []int{1, 2, 3, 4, 5}
	fmt.Println(FindSubSumKadane(testCase3))

	testCase4 := []int{-4, -3, -6, -10}
	fmt.Println(FindSubSumKadane(testCase4))
}
