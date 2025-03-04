package main

func DeleteDuplicate(numbers []int) []int {
	var result []int
	set := make(map[int]struct{})
	for _, numb := range numbers {
		_, ok := set[numb]
		if !ok {
			result = append(result, numb)
			set[numb] = struct{}{}
		}
	}
	return result
}
