package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "fmt"

// Приведём пример с фабрикой и созданием продуктов

// Интерфейс для создания продуктов
type Product interface {
	GetName() string
}

// Продукты
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "Product A"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "Product B"
}

// Интерфейс для фабрики
type Factory interface {
	CreateProduct() Product
}

// Фабрики
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	fmt.Println("Product from Factory A:", productA.GetName())

	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	fmt.Println("Product from Factory B:", productB.GetName())
}

/*
Паттерн "Фабричный метод" - это порождающий паттерн проектирования, который предоставляет
интерфейс для создания объектов, но позволяет подклассам выбирать класс создаваемых объектов.

Плюсы паттерна:
1. Гибкость и расширяемость кода.
2. Уменьшение зависимостей между классами.
3. Улучшенная поддержка кода и простота тестирования.

Минусы паттерна:
1. Увеличение количества классов в системе, что может усложнить структуру.
2. Возможность создания слишком многих вариантов объектов что усложнить понимание работы.
*/
