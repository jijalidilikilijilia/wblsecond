package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов командной строки
	column := flag.Int("k", 0, "Номер колонки для сортировки (по умолчанию 0, разделитель - пробел)")
	numeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	// Открытие файла на чтение
	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Считывание строк из файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Функция для сравнения строк при сортировке
	comparator := func(i, j int) bool {
		a := lines[i]
		b := lines[j]

		// Разделение строк на колонки
		aColumns := splitColumns(a)
		bColumns := splitColumns(b)

		// Выбор колонки для сортировки (пересчитываем номер колонки)
		columnIndex := *column - 1
		if columnIndex >= 0 && columnIndex < len(aColumns) && columnIndex < len(bColumns) {
			a = aColumns[columnIndex]
			b = bColumns[columnIndex]
		}

		// Сравнение строк
		if *numeric {
			// Сортировка по числовому значению
			na, errA := strconv.Atoi(a)
			nb, errB := strconv.Atoi(b)
			if errA == nil && errB == nil {
				return na < nb
			}
		}

		return a < b
	}

	// Выполнение сортировки
	if *unique {
		// Удаление повторяющихся строк
		lines = removeDuplicates(lines)
	}

	// Выполнение сортировки в соответствии с флагами
	if *reverse {
		sort.SliceStable(lines, func(i, j int) bool {
			return comparator(j, i) // Инвертируем порядок для сортировки в обратном порядке
		})
	} else {
		sort.SliceStable(lines, comparator)
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Функция для разделения строки на колонки (по пробелу)
func splitColumns(line string) []string {
	return strings.Fields(line)
}

// Функция для удаления повторяющихся строк
func removeDuplicates(lines []string) []string {
	var res []string
	uniqueLines := make(map[string]struct{})
	for _, elem := range lines {
		if _, ok := uniqueLines[elem]; !ok {
			uniqueLines[elem] = struct{}{}
			res = append(res, elem)
		}
	}
	return res
}
