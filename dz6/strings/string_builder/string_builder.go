package stringbuilder

import (
	"fmt"
	"strings"
	"time"
)

// StringBuilder свой тип, который позволяет собирать строку эффективнее
type StringBuilder struct {
	data []byte
}

// Write записывает len(p) байт в data
func (sb *StringBuilder) Write(p []byte) (n int, err error) {
	sb.data = append(sb.data, p...)
	return len(p), err
}

// Write записывает string в data
func (sb *StringBuilder) WriteString(s string) (n int, err error) {
	return sb.Write([]byte(s))
}

// String метод для определения вывода класса
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
