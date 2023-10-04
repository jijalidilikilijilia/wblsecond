package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		// Выводим приглашение для пользователя
		fmt.Print("$ ")

		// Считываем ввод пользователя
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при чтении ввода:", err)
			continue
		}

		// Удаляем символ новой строки из ввода
		input = strings.TrimSpace(input)

		// Разбиваем ввод на аргументы
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		// Обработка команд
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "Не указана целевая директория")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при смене директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при получении текущей директории:", err)
			} else {
				fmt.Println(dir)
			}
		case "echo":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "Нет аргументов для вывода")
				continue
			}
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "Не указан процесс для завершения")
				continue
			}
			pid := args[1]
			cmd := exec.Command("taskkill", "/F", "/T", "/PID", pid)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при завершении процесса:", err)
			}
		case "ps":
			cmd := exec.Command("tasklist")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды tasklist:", err)
			}
		case "exec":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "Не указана команда для выполнения")
				continue
			}
			cmd := exec.Command(args[1], args[2:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды:", err)
			}
		default:
			fmt.Fprintln(os.Stderr, "Неизвестная команда:", args[0])
		}
	}
}
