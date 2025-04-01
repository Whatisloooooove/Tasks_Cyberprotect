package main

import (
	"fmt"
	"sync"
)

type MyChannel struct {
	buf    []int      // Буфер для хранения данных
	mutex  sync.Mutex // Мьютекс для синхронизации доступа
	cond   *sync.Cond // Условная переменная для уведомлений
	closed bool       // Флаг закрытия канала
	cap    int        // Вместимость буфера
}

// Send отправляет данные в канал
func (mc *MyChannel) Send(value int) error {
    mc.mutex.Lock()
	defer mc.mutex.Unlock()

	if mc.closed {
		return fmt.Errorf("канал уже закрыт")
	}

	for len(mc.buf) == mc.cap {
		mc.cond.Wait()
		if mc.closed {
			return fmt.Errorf("канал уже закрыт")
		}	
	}

	mc.buf = append(mc.buf, value)
	mc.cond.Broadcast()
	return nil
}

// Receive достаёт значение из канала и возвращает вместе с булевым параметром(удалось ли прочитать)
func (mc *MyChannel) Receive() (int, bool) {
    mc.mutex.Lock()
	defer mc.mutex.Unlock()

	for len(mc.buf) == 0 {
		if mc.closed {
			return 0, false
		}
		mc.cond.Wait()
	}

	val := mc.buf[0]
	mc.buf = mc.buf[1:]
	mc.cond.Broadcast()
	return val, true
}

// Close закрывает канал
func (mc *MyChannel) Close() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	mc.closed = true
	mc.cond.Broadcast()
}
