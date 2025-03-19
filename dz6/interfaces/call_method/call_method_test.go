package callmethod

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

var ErrorNoMethod = errors.New("нет такого метода")

// Programist класс для реализации работы пользователя со спринтами,
// также хранит слайс знаий пользователя
type Programist struct {
	Name            string
	StackKnowledges []string
	SprintTasks     []string
}

var (
	tasks = []string{"настроить CI", "настроить CD"}
	tech  = []string{"CI", "CD"}
	date  = time.Date(2025, 3, 25, 0, 0, 0, 0, time.UTC)
)

// GetTasks позволяет пользователю получить задачи и время их исполнения
func (p *Programist) GetTasks(date time.Time, tech []string, tasks []string) string {
	if time.Until(date).Hours()/24 > 3 {
		p.StackKnowledges = append(p.StackKnowledges, tech...)
		p.SprintTasks = append(p.SprintTasks, tasks...)
		return ""
	}
	return "Я не смогу сделать так много("
}

// Тест на вызов CallMethod в случае, когда метод у класса есть
func TestCallMethodNoError(t *testing.T) {
	prog := &Programist{
		Name:            "Влад",
		StackKnowledges: []string{"C++", "Go"},
		SprintTasks:     []string{},
	}

	expected := &Programist{
		Name:            "Влад",
		StackKnowledges: []string{"C++", "Go", "CI", "CD"},
		SprintTasks:     []string{"настроить CI", "настроить CD"},
	}

	result, err := CallMethod(prog, "GetTasks", date, tech, tasks)

	require.NoError(t, err, "Ожидалось, что метод будет найден и вызван без ошибок")

	require.True(t, cmp.Equal(expected, prog), "После вызова метод должен изменить объект")

	require.Len(t, result, 1, "Ожидался один возвращаемый результат")
	require.Equal(t, "", result[0].Interface(), "Ожидалось, что метод вернет пустую строку")
}

// Тест на вызов CallMethod в случае, когда метод у класса нет
func TestCallMethodNoMethodError(t *testing.T) {
	prog := &Programist{Name: "Влад"}

	_, err := CallMethod(prog, "NonExistentMethod")

	require.Error(t, err, "Ожидалась ошибка при вызове несуществующего метода")
	require.Equal(t, ErrorNoMethod.Error(), err.Error(), "Ошибка должна соответствовать NoMethodError")
}
