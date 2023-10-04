package patter

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
import "fmt"

// Приведём пример созданием геометрических фигур

// Интерфейс для геометрических фигур
type Shape interface {
	Accept(Visitor)
}

// Реализация круга
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

// Реализация квадрата
type Square struct {
	Side float64
}

func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}

// Интерфейс для посетителя
type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
}

// Конкретный посетитель: Расчет площади
type AreaVisitor struct {
	TotalArea float64
}

func (av *AreaVisitor) VisitCircle(c *Circle) {
	av.TotalArea += 3.14 * c.Radius * c.Radius
}

func (av *AreaVisitor) VisitSquare(s *Square) {
	av.TotalArea += s.Side * s.Side
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Square{Side: 4},
		&Circle{Radius: 3},
	}

	// Даём посетителю возможность высчитать площадь, не меняя код самих фигур
	areaVisitor := &AreaVisitor{}
	for _, shape := range shapes {
		shape.Accept(areaVisitor)
	}

	fmt.Printf("Общая площадь всех фигур: %.2f\n", areaVisitor.TotalArea)
}

// Паттерн "Посетитель" позволяет добавлять новые операции к объектам без изменения их классов.
// Плюсы:
// 1. Разделение алгоритма от структуры объектов, что улучшает расширяемость кода.
// 2. Упрощение добавления новых операций к существующим классам объектов.
// 3. Позволяет работать с разнообразными классами объектов, не заботясь о их типах.
//
// Минусы:
// 1. Может привести к увеличению числа классов и интерфейсов в системе, что усложняет структуру.
// 2. Не всегда подходит для простых операций и может усложнить код в таких случаях.
