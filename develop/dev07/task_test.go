package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	// Создаем каналы для имитации сигналов.
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	// Создаем канал для сигнализации о завершении горутин.
	done := make(chan struct{})

	// Запускаем функцию or в горутине и передаем ей созданные каналы.
	go func() {
		defer close(done)
		<-or(ch1, ch2, ch3)
	}()

	// Закрываем один из каналов после задержки.
	time.Sleep(100 * time.Millisecond)
	close(ch1)

	// Задаем еще некоторую задержку, чтобы обеспечить выполнение остальной логики.
	time.Sleep(100 * time.Millisecond)

	// Ожидаем завершения всех горутин.
	<-done
}

func TestSig(t *testing.T) {
	// Создаем сигнальный канал и передаем его в функцию sig.
	signal := sig(100 * time.Millisecond)

	// Ждем, пока канал не будет закрыт.
	<-signal

	// Если канал не был закрыт, тест неудачен.
	if _, ok := <-signal; ok {
		t.Error("Expected the channel to be closed, but it's still open")
	}
}

func TestMainFunction(t *testing.T) {
	// Тестирование функции main, которая использует функцию or.
	// Здесь вы можете использовать CaptureOutput или другие способы проверки вывода.
}
