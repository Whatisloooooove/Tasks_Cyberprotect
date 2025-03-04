package main

func IsPalindrom(runes []rune) bool {
	leftInd, rightInd := 0, len(runes)-1
	for leftInd <= rightInd {
		if runes[leftInd] != runes[rightInd] {
			return false
		}
		leftInd++
		rightInd--
	}
	return true
}
