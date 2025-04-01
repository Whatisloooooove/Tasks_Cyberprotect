package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

var (
	ErrTimeout  = errors.New("таймаут записи")
	ErrCanceled = errors.New("context canceled")
)

// SafeSend безопасно отправлят данные в канал ch
func SafeSend(ctx context.Context, ch chan<- int, value int) error {
	select {
	case <-ctx.Done():
		return ErrCanceled
	case ch <- value:
		return nil
	case <-time.After(1 * time.Second):
		return ErrTimeout
	}
}

func main() {
	ch1 := make(chan int, 1)
	fmt.Printf("Канал: len=%d, cap=%d\n", len(ch1), cap(ch1))
	err1 := SafeSend(context.Background(), ch1, 42)
	log.Println("Сценарий 1:", err1)

	ch2 := make(chan int)
	fmt.Printf("Канал: len=%d, cap=%d\n", len(ch1), cap(ch1))
	err2 := SafeSend(context.Background(), ch2, 42)
	log.Println("Сценарий 2:", err2)

	ctx3, cancel3 := context.WithCancel(context.Background())
	cancel3()
	fmt.Printf("Канал: len=%d, cap=%d\n", len(ch1), cap(ch1))
	err3 := SafeSend(ctx3, ch1, 42)
	log.Println("Сценарий 3:", err3)
}
