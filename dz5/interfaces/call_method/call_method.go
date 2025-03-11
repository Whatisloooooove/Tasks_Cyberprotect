package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func CallMethod(obj interface{}, methodName string, args ...interface{}) (result []reflect.Value, err error) {
	reflVal := reflect.ValueOf(obj)
	method := reflVal.MethodByName(methodName)
	if !method.IsValid() {
		return nil, errors.New("нет такого метода")
	}

	reflArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflArgs[i] = reflect.ValueOf(arg)
	}

	return method.Call(reflArgs), nil
}

type Programist struct {
	Name            string
	StackKnowledges []string
	SprintTasks     []string
}

func (p *Programist) GetTasks(date time.Time, tech []string, tasks []string) string {
	now := time.Now()
	diff := now.Sub(date)
	if diff/24 > 3 {
		p.StackKnowledges = append(p.StackKnowledges, tech...)
		p.SprintTasks = append(p.SprintTasks, tasks...)
		return ""
	}
	return "Я не смогу сделать так много("
}

func (p *Programist) ShowTasks() {
	fmt.Println("Текущие задачи:", p.SprintTasks)
}

func main() {
	prog := &Programist{
		Name:            "Влад",
		StackKnowledges: []string{"C++", "Go"},
		SprintTasks:     []string{},
	}

	_, err := CallMethod(prog, "ShowTasks")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println(prog)

	tasks := []string{"настроить CI", "настроить CD"}
	tech := []string{"CI", "CD"}
	date := time.Date(2025, 16, 3, 0, 0, 0, 0, time.UTC)

	ans, err := CallMethod(prog, "GetTasks", date, tech, tasks)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	if len(ans) != 0 {
		fmt.Println(ans[0].Interface().(string))
	} else {
		fmt.Println(prog)
	}
}
