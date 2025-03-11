package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type ByName []User
type ByAge []User

func (n ByName) Len() int           { return len(n) }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	users := []User{
		{"Лена", 25},
		{"Влад", 20},
		{"Олег", 42},
		{"Маша", 30},
		{"Никита", 13},
	}

	sort.Sort(ByName(users))
	fmt.Println("Сортировка по именам: ", users)

	sort.Sort(ByAge(users))
	fmt.Println("Сортировка по возратсу: ", users)
}
