package processvalue

import (
	"fmt"
	"strings"
)

// ProcessValue определяет тип переданного интерфейса и 
// выводит в верхнем регистре, если это строка, 
// умножает на два, если это числовой тип
func ProcessValue(v interface{}) {
	switch val := v.(type) {
	case string:
		fmt.Println(strings.ToUpper(val))

	case int:
		fmt.Println(val * 2)
	case int8:
		fmt.Println(val * 2)
	case int16:
		fmt.Println(val * 2)
	case int32:
		fmt.Println(val * 2)
	case int64:
		fmt.Println(val * 2)

	case uint:
		fmt.Println(val * 2)
	case uint8:
		fmt.Println(val * 2)
	case uint16:
		fmt.Println(val * 2)
	case uint32:
		fmt.Println(val * 2)
	case uint64:
		fmt.Println(val * 2)

	case float32:
		fmt.Println(val * 2)
	case float64:
		fmt.Println(val * 2)

	case complex64:
		fmt.Println(val * 2)
	case complex128:
		fmt.Println(val * 2)
	default:
		fmt.Println("Неизвестный тип")
	}

}
