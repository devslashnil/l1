// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

package task

import "math"

// Point c инкапсулированными x,y
// маленькая буква означает недоступность не из пакета task
type Point struct {
	x, y float64
}

// NewPoint конструктор
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance считает Eвклидову метрику в двумерии
func (p *Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}
