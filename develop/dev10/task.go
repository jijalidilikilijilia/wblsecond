package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Определение флага для таймаута
	var timeoutStr string
	flag.StringVar(&timeoutStr, "timeout", "10s", "Таймаут для подключения (например, '3s')")
	flag.Parse()

	// Получение хоста и порта из аргументов командной строки
	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("Использование: go-telnet --timeout=3s host port")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	// Парсинг таймаута
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		fmt.Printf("Ошибка парсинга таймаута: %v\n", err)
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Printf("Ошибка подключения: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Подключено к", address)
	fmt.Println("Введите текст и нажмите Enter (Ctrl+D для выхода):")

	// Создаем канал для ожидания завершения ввода с клавиатуры
	done := make(chan struct{})
	go func() {
		for {
			data := make([]byte, 1024)
			n, err := os.Stdin.Read(data)
			if err != nil {
				close(done)
				return
			}
			conn.Write(data[:n])
		}
	}()

	// Читаем данные из сокета и выводим их в STDOUT
	go func() {
		for {
			data := make([]byte, 1024)
			n, err := conn.Read(data)
			if err != nil {
				if err == io.EOF {
					fmt.Println("\nСоединение закрыто сервером")
				} else {
					fmt.Printf("Ошибка чтения из сокета: %v\n", err)
				}
				close(done)
				return
			}
			fmt.Print(string(data[:n]))
		}
	}()

	// Ожидаем завершения ввода или сигнала завершения
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	select {
	case <-done:
	case <-c:
		fmt.Println("\nСоединение завершено.")
	}
}
