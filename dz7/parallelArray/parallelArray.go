package main

import (
	"fmt"
	"sync"
)

func ParallelSumArray(array []int, numGoroutines int) int {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	sum := 0
	for i := 0; i < numGoroutines; i++ {
		start := i * len(array) / numGoroutines
		end := (i + 1) * len(array) / numGoroutines
		if i == numGoroutines-1 {
			end = len(array)
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localSum := 0
			for _, v := range array[start:end] {
				localSum += v
			}
			mutex.Lock()
			sum += localSum
			mutex.Unlock()
		}(start, end)
	}
	wg.Wait()
	return sum
}

func main() {
	array := make([]int, 100)
	for i := 0; i < 100; i++ {
		array[i] = i + 1
	}

	sum := ParallelSumArray(array, 100)
	fmt.Println(sum)
}
