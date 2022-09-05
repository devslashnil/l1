// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

package task

import (
	"fmt"
)

// GetTypeOf возвращает %T a Go-syntax representation of the type of the value
func GetTypeOf(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
