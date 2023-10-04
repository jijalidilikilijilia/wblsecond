package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

// Интерфейс для состояний
type State interface {
	InsertMoney()
	ChooseDrink()
	Dispense()
}

// Состояние "Ожидание внесения денег"
type WaitingForMoneyState struct{}

func (s *WaitingForMoneyState) InsertMoney() {
	fmt.Println("Деньги внесены. Выберите напиток.")
}

func (s *WaitingForMoneyState) ChooseDrink() {
	fmt.Println("Пожалуйста, внесите деньги сначала.")
}

func (s *WaitingForMoneyState) Dispense() {
	fmt.Println("Пожалуйста, внесите деньги и выберите напиток.")
}

// Состояние "Выбор напитка"
type ChoosingDrinkState struct{}

func (s *ChoosingDrinkState) InsertMoney() {
	fmt.Println("Деньги уже внесены. Выберите напиток.")
}

func (s *ChoosingDrinkState) ChooseDrink() {
	fmt.Println("Напиток выбран. Ожидайте выдачи.")
}

func (s *ChoosingDrinkState) Dispense() {
	fmt.Println("Пожалуйста, сначала выберите напиток.")
}

// Автомат для продажи напитков
type BeverageMachine struct {
	state State
}

func (m *BeverageMachine) setState(newState State) {
	m.state = newState
}

func (m *BeverageMachine) InsertMoney() {
	m.state.InsertMoney()
}

func (m *BeverageMachine) ChooseDrink() {
	m.state.ChooseDrink()
}

func (m *BeverageMachine) Dispense() {
	m.state.Dispense()
}

func main() {
	machine := &BeverageMachine{}
	machine.setState(&WaitingForMoneyState{})

	machine.InsertMoney() // Деньги внесены. Выберите напиток.
	machine.ChooseDrink() // Пожалуйста, внесите деньги сначала.
	machine.Dispense()    // Пожалуйста, внесите деньги и выберите напиток.

	machine.setState(&ChoosingDrinkState{})
	machine.InsertMoney() // Деньги уже внесены. Выберите напиток.
	machine.ChooseDrink() // Напиток выбран. Ожидайте выдачи.
	machine.Dispense()    // Пожалуйста, сначала выберите напиток.

	// Паттерн "Состояние" используется здесь для управления состояниями автомата для продажи напитков.
	// Плюсы паттерна:
	// 1. Изолирование состояний: Каждое состояние инкапсулируется в отдельном классе, что упрощает поддержку и добавление новых состояний.
	// 2. Уменьшение условной логики: Паттерн позволяет избежать множества условных операторов, что делает код более чистым и читаемым.
	// 3. Расширяемость: Вы можете легко добавлять новые состояния и изменять поведение объекта без изменения существующего кода.

	// Минусы паттерна:
	// 1. Увеличение числа классов: Использование этого паттерна может привести к созданию большого количества классов, особенно если у вас есть много состояний.
	// 2. Сложность: Для простых случаев добавление паттерна "Состояние" может быть избыточным и увеличивать сложность кода.
}
