package main

func FindSubSumKadane(numbers []int) int {
	current_sum, max_sum := numbers[0], numbers[0]
	for i := 1; i < len(numbers); i++ {
		current_sum = max(numbers[i], numbers[i]+current_sum)
		max_sum = max(max_sum, current_sum)
	}
	return max_sum
}
