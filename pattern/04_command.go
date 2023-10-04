package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
import "fmt"

// Приведём пример с работой светильника

// Наш receiver который получает команды
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Светильник включен")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Светильник выключен")
}

// Command - интерфейс команды
type Command interface {
	Execute()
}

type TurnOnCommand struct {
	light *Light
}

func NewTurnOnCommand(light *Light) Command {
	return &TurnOnCommand{light}
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

type TurnOffCommand struct {
	light *Light
}

func NewTurnOffCommand(light *Light) Command {
	return &TurnOffCommand{light}
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// Invoker - инициатор команд
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	// Создаем светильник и команды для включения и выключения
	livingRoomLight := &Light{}
	turnOnCommand := NewTurnOnCommand(livingRoomLight)
	turnOffCommand := NewTurnOffCommand(livingRoomLight)

	// Создаем пульт управления и выполняем команды
	remoteControl := &RemoteControl{}
	remoteControl.SetCommand(turnOnCommand)

	// Нажимаем кнопку на пульте и включаем светильник
	remoteControl.PressButton()

	// Теперь переключаем команду на выключение и нажимаем кнопку
	remoteControl.SetCommand(turnOffCommand)
	remoteControl.PressButton()
}

/*
   Паттерн "Команда" - это поведенческий паттерн, который инкапсулирует запросы в объекты,
   позволяет параметризовать клиентские объекты операциями, ставить запросы в очередь,
   поддерживать отмену операций и управлять последовательностью выполнения операций.

   Плюсы:
   1. Изолирует отправителя и получателя запроса.
   2. Позволяет легко добавлять новые команды и изменять их параметры.
   3. Поддерживает отмену и повторение операций.
   4. Упрощает управление очередью операций и планированием.

   Минусы:
   1. Может привести к созданию множества классов команд, если операций много.
   2. Увеличивает сложность кода из-за необходимости создания дополнительных классов.
*/
