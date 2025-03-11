package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	result := ""
	for i := 0; i < 1000; i++ {
		result += "a"
	}
	fmt.Println("+= time:", time.Since(start))

	start = time.Now()
	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString("a")
	}
	fmt.Println("Builder time:", time.Since(start))
}
