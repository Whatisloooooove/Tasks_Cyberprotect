package main

import (
	"fmt"
	"math/rand"
)

//-------------- Input --------------------

func GetUserAttempt(bord int) int {
	num := 0
	for num == 0 {
		fmt.Printf("Попробуйте угадать загаданное число от 1 до %d:\n", bord)
		_, err := fmt.Scan(&num)
		if err != nil || num < 1 || num > bord {
			fmt.Println("Введите корректное число!")
		}
	}
	return num
}

func SelectDifficult() int {
	fmt.Println("Выберите уровень сложности:")
	fmt.Println("Лёгкий")
	fmt.Println("Средний")
	fmt.Println("Сложный")

	var difficult string
	bord := 0
	for bord == 0 {
		fmt.Scan(&difficult)
		switch difficult {
		case "Лёгкий":
			bord = 50
		case "Средний":
			bord = 100
		case "Сложный":
			bord = 500
		default:
			fmt.Println("Введите корректный уровень сложности из предложенных выше")
		}
	}
	fmt.Println("Вы выбрали режим:", difficult)

	return bord
}

//-------------- UseCase ------------------

func RandomNum(max int) int {
	return rand.Intn(max)*21092004%max + 1
}

func main() {
	max_bord := SelectDifficult()
	goal_num := RandomNum(max_bord)

	for counter_attemps := 1; ; counter_attemps++ {
		answer_num := GetUserAttempt(max_bord)
		if answer_num > goal_num {
			fmt.Println("Слишком много!")
		} else if answer_num < goal_num {
			fmt.Println("Слишком мало!")
		} else {
			fmt.Printf("Поздравляем! Вы угадали число за %d попыток\n", counter_attemps)
			break
		}
	}
}
