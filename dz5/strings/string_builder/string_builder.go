package main

import (
	"fmt"
	"strings"
	"time"
)

type StringBuilder struct {
	data []byte
}

func (sb *StringBuilder) Write(p []byte) (n int, err error) {
	sb.data = append(sb.data, p...)
	return len(p), err
}

func (sb *StringBuilder) WriteString(s string) (n int, err error) {
	return sb.Write([]byte(s))
}

func (sb *StringBuilder) String() string {
	return string(sb.data)
}

func main() {
	start := time.Now()
	var builder strings.Builder
	for i := 0; i < 1; i++ {
		builder.WriteString("a")
	}
	fmt.Println("Strings builder time:", time.Since(start))

	start = time.Now()
	var myBuilder StringBuilder
	for i := 0; i < 1; i++ {
		myBuilder.WriteString("a")
	}
	fmt.Println("My builder time:", time.Since(start))

}

// По времени StringBuilder оказался более эффективным чем strings.Builder
// (но почему-то при повторном запуске с одними и теми же значениями выдавало различное время
// и иногда strings.Builder выдавал время меньше чем в StringBuilder
