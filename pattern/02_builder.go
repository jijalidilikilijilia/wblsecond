package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
import "fmt"

//Приведём пример с созданием пиццы

// Определение структуры пиццы
type Pizza struct {
	Size     string
	Crust    string
	Toppings []string
}

// Определение интерфейса "Строителя"
type PizzaBuilder interface {
	SetSize(size string)
	SetCrust(crust string)
	AddTopping(topping string)
	BuildPizza() *Pizza
}

// Создание структуры "Строитель пиццы"
type ConcretePizzaBuilder struct {
	pizza *Pizza
}

func NewPizzaBuilder() *ConcretePizzaBuilder {
	return &ConcretePizzaBuilder{pizza: &Pizza{}}
}

func (pb *ConcretePizzaBuilder) SetSize(size string) {
	pb.pizza.Size = size
}

func (pb *ConcretePizzaBuilder) SetCrust(crust string) {
	pb.pizza.Crust = crust
}

func (pb *ConcretePizzaBuilder) AddTopping(topping string) {
	pb.pizza.Toppings = append(pb.pizza.Toppings, topping)
}

func (pb *ConcretePizzaBuilder) BuildPizza() *Pizza {
	return pb.pizza
}

// Пример работы
func main() {
	builder := NewPizzaBuilder()

	// Создаем пиццу с разными характеристиками
	builder.SetSize("Средний")
	builder.SetCrust("Тонкае")
	builder.AddTopping("Ананас")
	builder.AddTopping("Пеперони")

	pizza := builder.BuildPizza()

	// Выводим информацию о пицце
	fmt.Println("Размер:", pizza.Size)
	fmt.Println("Корочка:", pizza.Crust)
	fmt.Println("Начинка:", pizza.Toppings)
}

// Паттерн Строитель представляет собой способ создания сложных объектов,
// разбивая процесс инициализации на отдельные шаги. Это упрощает создание объектов
// с разными конфигурациями и делает код более читаемым и поддерживаемым.

// Плюсы паттерна Строитель:
// 1. Упрощает создание сложных объектов.
// 2. Позволяет создавать объекты с разными конфигурациями.
// 3. Избегает длинных списков аргументов в конструкторах.
// 4. Улучшает читаемость кода и его поддерживаемость.

// Минусы паттерна Строитель:
// 1. Требует создания дополнительных классов и интерфейсов, что может усложнить код.
// 2. В некоторых случаях может привести к увеличению объема кода.
