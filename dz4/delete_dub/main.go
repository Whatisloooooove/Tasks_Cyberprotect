package main

import (
	"fmt"
)

func main() {
	testCase1 := []int{}
	fmt.Println(DeleteDuplicate(testCase1))

	testCase2 := []int{1, 1, 1}
	fmt.Println(DeleteDuplicate(testCase2))

	testCase3 := []int{1, 1, 1, 2, 2, 4, 3, 2, 1, 2, 3}
	fmt.Println(DeleteDuplicate(testCase3))
}
