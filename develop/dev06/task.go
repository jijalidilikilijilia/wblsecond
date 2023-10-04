package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов
	delimiterFlag := flag.String("d", "\t", "Указать разделитель полей (по умолчанию TAB)")
	fieldsFlag := flag.String("f", "", "Выбрать поля/колонки)")
	separatedFlag := flag.Bool("s", false, "Вывести только строки с разделителем")
	flag.Parse()

	// Открытие файла для чтения
	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создание сканера для чтения строк из файла
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Проверка, содержит ли строка разделитель
		if strings.Contains(line, *delimiterFlag) || !*separatedFlag {
			// Разделение строки на поля с использованием разделителя
			fields := strings.Split(line, *delimiterFlag)

			// Вывод запрошенных полей
			if *fieldsFlag != "" {
				fieldIndices := strings.Split(*fieldsFlag, ",")
				outputFields := make([]string, len(fieldIndices))
				for i, idxStr := range fieldIndices {
					idx := parseFieldIndex(idxStr)
					if idx >= 0 && idx < len(fields) {
						outputFields[i] = fields[idx]
					}
				}
				fmt.Println(strings.Join(outputFields, *delimiterFlag))
			} else {
				// Если не указан флаг -f то выводим всю строку
				fmt.Println(line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}

// Функция для преобразования индекса поля из строки в число
func parseFieldIndex(fieldIdxStr string) int {
	if idx, err := strconv.Atoi(fieldIdxStr); err == nil {
		return idx - 1
	}
	return -1
}

// go run main.go -f 1,2 data.txt
//
