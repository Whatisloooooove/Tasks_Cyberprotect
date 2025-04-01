package main

import (
	"context"
	"fmt"
	"time"
)

func Fibonacci(ctx context.Context, rateLimiter <-chan time.Time) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		a, b := 0, 1
		for {
			select {
			case <-ctx.Done():
				return
			case <-rateLimiter:
				select {
				case ch <- a:
					a, b = b, a+b
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rateLimiter := time.Tick(time.Second)
	fibCh := Fibonacci(ctx, rateLimiter)

	for i := 0; i < 10; i++ {
		fmt.Println(<-fibCh)
	}
}
