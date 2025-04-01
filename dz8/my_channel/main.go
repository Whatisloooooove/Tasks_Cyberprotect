package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := &MyChannel{
		buf:    make([]int, 0, 2),
		closed: false,
		cap:    2,
	}
	ch.cond = sync.NewCond(&ch.mutex)

	go func() {
		for i := 0; i < 5; i++ {
			ch.Send(i)
			fmt.Printf("В канал отправлено число %d\n", i)
		}
		ch.Close()
	}()

	for {
		val, ok := ch.Receive()
		if !ok {
			break
		}
		fmt.Printf("Из канала прочитано число %d\n", val)
	}
}
