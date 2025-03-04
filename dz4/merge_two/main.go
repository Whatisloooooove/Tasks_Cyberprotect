package main

import (
	"fmt"
)

func main() {
	testCase1_1 := []int{}
	testCase1_2 := []int{0}
	fmt.Println(MergeTwoSlices(testCase1_1, testCase1_2))

	testCase2_1 := []int{1, 1, 1}
	testCase2_2 := []int{2, 2, 2}
	fmt.Println(MergeTwoSlices(testCase2_1, testCase2_2))

	testCase3_1 := []int{1, 1, 1, 2, 2, 4, 3, 2}
	testCase3_2 := []int{2, 5, 3, 4, 3}
	fmt.Println(MergeTwoSlices(testCase3_1, testCase3_2))
}
