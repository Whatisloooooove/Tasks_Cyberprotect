package sprintf

import "fmt"

// Person хранит информацию пользователя Name и Age
type Person struct {
	Name string
	Age  int
}

// String возвращает строковое представление структуры Person
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}
