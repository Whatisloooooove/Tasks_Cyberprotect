package main

import (
	"fmt"
	"os"
)

func safeFileWrite(filename string, content string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func main() {
	safeFileWrite("example.txt", "safe write")
}
