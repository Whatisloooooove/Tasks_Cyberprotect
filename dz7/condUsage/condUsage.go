package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	cond := sync.NewCond(&mutex)
	var ready bool

	wg.Add(1)
	go func() {
		defer wg.Done()
		mutex.Lock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Горутина получила сигнал!")
		mutex.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mutex.Lock()
		ready = true
		cond.Signal()
		mutex.Unlock()
	}()
	wg.Wait()
}
