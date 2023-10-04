package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определение флагов командной строки
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк перед совпадением")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")
	count := flag.Bool("c", false, "количество строк")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")

	// Парсинг флагов
	flag.Parse()

	// Получение аргумента для поиска
	searchArg := flag.Arg(0)
	if searchArg == "" {
		fmt.Println("Необходимо указать строку для поиска.")
		return
	}

	// Открытие файла для чтения
	fileName := flag.Arg(1)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close()

	// Создание сканера для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Переменные для хранения найденных строк и номера строки
	var matchingLines []string
	var lineNumber int

	// Функция для проверки совпадения строки
	match := func(line string) bool {
		// Используем strings.Contains для флага -v с учётом флага -i
		if *ignoreCase {
			return strings.Contains(strings.ToLower(line), strings.ToLower(searchArg)) != *invert
		}
		if *fixed {
			return line == searchArg != *invert
		}
		return strings.Contains(line, searchArg) != *invert
	}

	// Переменная для хранения строк перед совпадением
	var beforeLines []string

	// Перебор строк в файле
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		// Обработка флага -C (печать ±N строк вокруг совпадения)
		if match(line) {
			// Добавление строк перед совпадением
			for _, bline := range beforeLines {
				if *lineNum {
					bline = fmt.Sprintf("%d: %s", lineNumber-len(beforeLines), bline)
				}
				matchingLines = append(matchingLines, bline)
			}

			// Добавление строки с совпадением в результаты с учётом флага -n
			if *lineNum {
				line = fmt.Sprintf("%d: %s", lineNumber, line)
			}
			matchingLines = append(matchingLines, line)

			// Обработка флага -A (печать N строк после совпадения)
			for i := 1; i <= *after && scanner.Scan(); i++ {
				line = scanner.Text()
				lineNumber++
				if *lineNum {
					line = fmt.Sprintf("%d: %s", lineNumber, line)
				}
				matchingLines = append(matchingLines, line)
			}

			// Обработка флага -C (печать ±N строк вокруг совпадения)
			for i := 1; i <= *context && scanner.Scan(); i++ {
				line = scanner.Text()
				lineNumber++
				if *lineNum {
					line = fmt.Sprintf("%d: %s", lineNumber, line)
				}
				matchingLines = append(matchingLines, line)
			}

			// Сброс буфера строк перед совпадением
			beforeLines = []string{}
		} else if *before > 0 || *context > 0 {
			// Добавление строк перед совпадением в буфер
			beforeLines = append(beforeLines, line)

			// Ограничение количества строк перед совпадением
			if len(beforeLines) > *before+*context {
				beforeLines = beforeLines[len(beforeLines)-(*before+*context):]
			}
		}
	}

	// Вывод результата в зависимости от флагов
	if *count {
		fmt.Printf("Количество совпадений: %d\n", len(matchingLines))
	} else if *invert {
		for _, line := range matchingLines {
			fmt.Println(line)
		}
	} else {
		for _, line := range matchingLines {
			fmt.Println(line)
		}
	}
}
