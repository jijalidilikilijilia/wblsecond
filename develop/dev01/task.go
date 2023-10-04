package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func getNTPTime() (time.Time, error) {
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}

func main() {
	// Получение точного времени с использования NTP
	ntpTime, err := getNTPTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %v\n", err)
		os.Exit(1)
	}

	// Вывод точного времени
	fmt.Printf("Текущее точное время: %v\n", ntpTime)

	// Пример использования времени в коде
	// например, вы можете получить час, минуту и секунду
	hour, minute, second := ntpTime.Clock()
	fmt.Printf("Час: %02d, Минута: %02d, Секунда: %02d\n", hour, minute, second)
}
