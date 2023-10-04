package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import "fmt"

// Приведём пример с выбором способа оплаты

// Интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Оплата кредитной картой
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f с помощью кредитной карты", amount)
}

// Оплата Qiwi
type QiwiPayment struct{}

func (p *QiwiPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f с помощью PayPal", amount)
}

// Контекст, использующий стратегию
type PaymentContext struct {
	PaymentStrategy PaymentStrategy
}

func (pc *PaymentContext) MakePayment(amount float64) string {
	return pc.PaymentStrategy.Pay(amount)
}

func main() {
	creditCardPayment := &CreditCardPayment{}
	qiwiPayment := &QiwiPayment{}

	paymentContext := &PaymentContext{}

	// Оплата с использованием кредитной карты
	paymentContext.PaymentStrategy = creditCardPayment
	fmt.Println(paymentContext.MakePayment(100.0))

	// Оплата с использованием
	paymentContext.PaymentStrategy = qiwiPayment
	fmt.Println(paymentContext.MakePayment(50.0))
}

/*
   Паттерн "Стратегия" - это поведенческий паттерн проектирования, который позволяет определить семейство
   алгоритмов, инкапсулировать каждый из них и обеспечить их взаимозаменяемость. Этот паттерн позволяет выбирать
   подходящий алгоритм на лету, в зависимости от конкретных условий, не изменяя код клиентского класса.

   Плюсы паттерна:
   1. Гибкость и расширяемость кода.
   2. Изоляция стратегий и уменьшение зависимостей.

   Минусы паттерна:
   1. Увеличение числа классов в системе, что может усложнить структуру.
*/
