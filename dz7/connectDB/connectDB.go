package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDB() {
	var once sync.Once
	var wg sync.WaitGroup

	initialize := func() {
		var err error
		db, err = sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
		if err != nil {
			fmt.Println("Ошибка подключения к БД:", err)
			return
		}
		fmt.Println("Подключение к БД выполнено")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}
	wg.Wait()
}

func main() {
	ConnectToDB()
}
