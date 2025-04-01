package main

import (
	"context"
	"fmt"
)

// FirstN вощврашает первые N элементов из канала
func FirstN(ctx context.Context, n int, input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		count := 0
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-input:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case output <- val:
					count++
					if count >= n {
						return
					}
				}
			}
		}
	}()

	return output
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := make(chan int)
	go func() {
		defer close(input)
		for i := 0; i < 10; i++ {
			input <- i
		}
	}()

	for val := range FirstN(ctx, 5, input) {
		fmt.Println(val)
	}
}
