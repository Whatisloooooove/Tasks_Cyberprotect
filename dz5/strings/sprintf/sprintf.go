package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	person := Person{Name: "Vlad", Age: 20}
	format := fmt.Sprintf("Name: %s, Age: %d", person.Name, person.Age)
	fmt.Println(format)

	fmt.Println(person)
}
