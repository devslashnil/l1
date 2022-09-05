// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

package task

import (
	"fmt"
	"reflect"
)

func GetTypeOf(i interface{}) string {
	return fmt.Sprintf("%s", reflect.TypeOf(i))
}

func GetTypeOf3(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan any:
		return "chan"
	}
	return ""
}
