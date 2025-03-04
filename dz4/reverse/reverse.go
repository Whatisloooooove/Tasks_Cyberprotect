package main

func Reverse(strs []string) []string {
	var result []string
	for i := len(strs) - 1; i >= 0; i-- {
		result = append(result, strs[i])
	}
	return result
}
