package main

import (
	"fmt"
	"sync"
)

func goroutineSum(left, right int, numGoroutines int) int {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	sum := 0
	for i := 0; i < numGoroutines; i++ {
		start := left + ((right-left)/numGoroutines)*i
		end := start + (right-left)/numGoroutines - 1

		if i == numGoroutines-1 {
			end = right
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localSum := 0
			for j := start; j <= end; j++ {
				localSum += j
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
	left := 11
	right := 1011
	sum := goroutineSum(left, right, 1000)

	expectedSum := 0
	for i := 11; i <= 1011; i++ {
		expectedSum += i
	}
	fmt.Println(expectedSum)
	fmt.Println(sum)
}
