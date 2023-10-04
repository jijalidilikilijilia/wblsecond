package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadSite(t *testing.T) {
	// Создаем временную директорию для тестовых файлов.
	tempDir := t.TempDir()

	// Заглушка для HTTP-сервера с использованием httptest
	server := mockHTTPServer(t, "Hello, World!")

	// Закрываем сервер после выполнения тестов
	defer server.Close()

	// Вызываем DownloadSite для загрузки сайта
	err := DownloadSite(server.URL, tempDir)
	if err != nil {
		t.Fatalf("Ошибка при скачивании сайта: %v", err)
	}

	// Проверяем, что файл index.html был создан в директории
	_, err = os.Stat(tempDir + "/index.html")
	if err != nil {
		t.Fatalf("Файл 'index.html' не был создан: %v", err)
	}

	// Читаем содержимое созданного файла
	fileContent, err := os.ReadFile(tempDir + "/index.html")
	if err != nil {
		t.Fatalf("Ошибка при чтении файла 'index.html': %v", err)
	}

	// Проверяем, что содержимое файла соответствует ожидаемой строке
	expectedContent := "Hello, World!"
	if string(fileContent) != expectedContent {
		t.Fatalf("Содержимое файла не соответствует ожидаемому. Ожидается: '%s', Фактический результат: '%s'", expectedContent, string(fileContent))
	}
}

// mockHTTPServer создает временный HTTP-сервер c содержимым
func mockHTTPServer(t *testing.T, response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(response))
		if err != nil {
			t.Fatalf("Ошибка при записи в ответ: %v", err)
		}
	}))
}
