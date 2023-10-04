package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadSite(url string, downloadDir string) error {
	// HTTP запрос к указанному URL
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Проверяем код ответа
	if response.StatusCode != http.StatusOK {
		return errors.New("StatusCode not ok")
	}

	// Создаем директорию для сохранения содержимого сайта, если она не существует
	if err := os.MkdirAll(downloadDir, os.ModePerm); err != nil {
		return err
	}

	// Создаем файл для сохранения сайта
	filePath := fmt.Sprintf("%s/index.html", downloadDir)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем содержимое сайта в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// URL сайта для скачивания
	url := "https://habr.com/ru/news/541958/"

	//Директория куда сохранять
	downloadDir := "."

	err := DownloadSite(url, downloadDir)
	if err != nil {
		fmt.Printf("Ошибка при скачивании сайта: %v\n", err)
		return
	}

	fmt.Printf("Сайт успешно скачан в директорию: %s\n", downloadDir)
}
