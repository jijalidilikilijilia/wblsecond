package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(input string) (string, error) {
	if input == "" {
		// Если строка пустая, возвращаем пустую строку
		return "", nil
	}

	str := []rune(input) // Преобразуем строку в слайс рун
	result := make([]rune, 0)

	if unicode.IsDigit(str[0]) {
		// Если строка начинается с цифры, считаем её некорректной и возвращаем ошибку.
		return "", errors.New("(некорректная строка)")
	}

	i := 0
	for i < len(str) {
		// Если найдено число
		if unicode.IsDigit(str[i]) {
			number := make([]rune, 0)

			// Считываем число из строки пока не упрёмся в след. букву
			for i < len(str) && unicode.IsDigit(str[i]) {
				number = append(number, str[i])
				i++
			}
			i--

			count, _ := strconv.Atoi(string(number)) // Преобразуем строку с числом в целое число
			for i := 0; i < count-1; i++ {
				// Повторяем предыдущий символ n-1 раз и добавляем к результату
				result = append(result, result[len(result)-1])
			}
		} else {
			// Если текущий символ не является цифрой, добавляем его к результату
			result = append(result, str[i])
		}
		i++
	}
	return string(result), nil // Преобразуем слайс рун обратно в строку и возвращаем результат
}

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
}
