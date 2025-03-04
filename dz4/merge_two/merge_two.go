package main

func MergeTwoSlices(slice1, slice2 []int) []int {
	set := make(map[int]struct{})
	result := []int{}

	for _, num := range slice1 {
		_, ok := set[num]
		if !ok {
			set[num] = struct{}{}
			result = append(result, num)
		}
	}

	for _, num := range slice2 {
		_, ok := set[num]
		if !ok {
			set[num] = struct{}{}
			result = append(result, num)
		}
	}

	return result
}
