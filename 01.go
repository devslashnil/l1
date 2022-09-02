// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от
// родительской структуры Human (аналог наследования).

package task

type Human struct{}

func (_ Human) Foo() {}

type Action struct {
	// применили встраивание
	Human
}
