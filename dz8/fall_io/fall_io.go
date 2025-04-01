package main

import (
	"fmt"
	"math/big"
	"math/rand"

	"sync"
	"time"
)

// Factorial считает фактоорил числа и возращает результат в виде *big.Int
func Factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func generateTasks(count int) []int {
	tasks := make([]int, count)
	for i := 0; i < count; i++ {
		tasks[i] = 10 + rand.Intn(11)
	}
	return tasks
}

// ProccessSequenially выполнят задачи в одной горутине
func ProcessSequentially(tasks []int) *big.Int {
	total := big.NewInt(0)
	for _, task := range tasks {
		total.Add(total, Factorial(task))
	}
	return total
}

func worker(tasks <-chan int, results chan<- *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range tasks {
		results <- Factorial(n)
	}
}

// ProcessConcurrently исполняет задачи в numWorkers горутинах
func ProcessConcurrently(tasks []int, numTasks, numWorkers int) *big.Int {
	tasksCh := make(chan int, numTasks)
	resultCh := make(chan *big.Int, numTasks)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(tasksCh, resultCh, &wg)
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()
	close(resultCh)

	total := big.NewInt(0)
	for res := range resultCh {
		total.Add(total, res)
	}
	return total
}

func main() {
	numWorkers := 10
	numTasks := 11

	tasks := generateTasks(numTasks)

	start := time.Now()
	totalSingle := ProcessSequentially(tasks)
	fmt.Printf("Однопоточная версия: %v, Сумма: %d\n", time.Since(start), totalSingle)

	start = time.Now()
	totalConcurrent := ProcessConcurrently(tasks, numTasks, numWorkers)
	fmt.Printf("Многопоточная: %v, Сумма: %d\n", time.Since(start), totalConcurrent)
}
