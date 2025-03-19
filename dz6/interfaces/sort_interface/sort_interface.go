package sortinterface

type User struct {
	Name string
	Age  int
}

// ByName реализует интерфейс sort.Interface для сортировки по имени
type ByName []User

func (n ByName) Len() int           { return len(n) }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

// ByAge реализует интерфейс sort.Interface для сортировки по возрасту
type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
