package main

import (
	"fmt"
)

func main() {
	testCase1 := []rune{'🙂', '🙂'}
	if IsPalindrom(testCase1) != true {
		fmt.Errorf("%v является палиндромом", testCase1)
	}

	testCase2 := []rune{'d', 'z', 'g', 't', 't'}
	if IsPalindrom(testCase2) == true {
		fmt.Errorf("%v не является палиндромом", testCase2)
	}

	testCase3 := []rune{'a', 'b', 'o', 'b', 'a'}
	if IsPalindrom(testCase3) != true {
		fmt.Errorf("%v является палиндромом", testCase3)
	}

	testCase4 := []rune{}
	if IsPalindrom(testCase4) != true {
		fmt.Errorf("%v является палиндромом", testCase4)
	}

}
