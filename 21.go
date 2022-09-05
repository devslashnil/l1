// Реализовать паттерн «адаптер» на любом примере.

package task

// Target интерфейс с которым должен уметь работать Adapter
type Target interface {
	Foo() int
}

// Adaptable структура для адаптирования
type Adaptable struct {
}

// NewAdapter конструктор для Adapter
func NewAdapter(a *Adaptable) Target {
	return &Adapter{a}
}

// Foo конкретная реализация
func (a *Adaptable) Foo() int {
	return 0
}

// Adapter имплементирует Target и является адамтером
type Adapter struct {
	*Adaptable
}

// Foo адаптивный метод
func (a *Adapter) Foo() int {
	// В этой секции можно сделать любые грязные дела
	// для этого у нужен адаптер
	return a.Foo()
}
