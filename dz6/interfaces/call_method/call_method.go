package callmethod

import (
	"errors"
	"reflect"
)

// CallMethod позволяет проверить наличие метода у класса
// и вызвать его, если он есть, от переданных параметров
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
