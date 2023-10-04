package pattern

/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Facade_pattern
*/

import "fmt"

// Приведём пример с запуском компьютера
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU: Запуск...")
}

type GPU struct{}

func (g *GPU) Start() {
	fmt.Println("GPU: Запуск...")
}

type Memory struct{}

func (m *Memory) Start() {
	fmt.Println("Память: Запуск...")
}

type Storage struct{}

func (s *Storage) Start() {
	fmt.Println("Хранилище: Запуск...")
}

// FacadeComputer предоставляет упрощенный интерфейс для запуска компьютера
type FacadeComputer struct {
	cpu     *CPU
	gpu     *GPU
	memory  *Memory
	storage *Storage
}

func NewFacadeComputer() *FacadeComputer {
	return &FacadeComputer{
		cpu:     &CPU{},
		gpu:     &GPU{},
		memory:  &Memory{},
		storage: &Storage{},
	}
}

// Start инициализирует и запускает компьютер с помощью фасада
func (cf *FacadeComputer) Start() {
	fmt.Println("ComputerFacade: Запуск компьютера...")
	cf.cpu.Start()
	cf.gpu.Start()
	cf.memory.Start()
	cf.storage.Start()
	fmt.Println("ComputerFacade: Компьютер запущен.")
}

func main() {
	// Создание экземпляра ComputerFacade
	computer := NewFacadeComputer()

	// Запуск компьютера с использованием фасада
	computer.Start()
}

// Паттерн "Фасад" предоставляет унифицированный интерфейс к набору интерфейсов в подсистеме,
// упрощая работу с ней и скрывая детали реализации

// Плюсы паттерна "Фасад":
// 1. Уменьшает сложность взаимодействия с подсистемой
// 2. Читаемость и обслуживаемость кода значительно повышается

// Минусы паттерна "Фасад":
// 1. Может привести к созданию больших и сложных фасадов
// 2. Изменнения подсистемы могут потребовать изменений фасада
