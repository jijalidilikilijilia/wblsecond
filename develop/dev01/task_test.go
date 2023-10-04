package main

import (
	"testing"
)

func TestGetNTPTime(t *testing.T) {
	t.Run("TestNTPTime", func(t *testing.T) {
		// Получение времени через функцию
		ntpTime, err := getNTPTime()
		if err != nil {
			t.Errorf("Ошибка при получении времени: %v", err)
		}

		// Проверка, что время не нулевое
		if ntpTime.IsZero() {
			t.Errorf("Получено нулевое время")
		}
	})
}
