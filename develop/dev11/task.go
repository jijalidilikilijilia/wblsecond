package main

// http://localhost:3000/events_for_day?date=2023-10-10&user_id=1
// http://localhost:3000/events_for_week?date=2023-10-10&user_id=1
// http://localhost:3000/events_for_month?date=2023-10-10&user_id=1

// http://localhost:3000/create_event
// {
//   "user_id": 1,
//   "title": "Первое событие 1 пользователь",
//   "date": "2023-10-20"
// }

// http://localhost:3000/update_event
// {
//   "event_id": 2,
//   "title": "Новое название события",
//   "date": "2023-10-09"
// }

// http://localhost:3000/delete_event
// {
//   "event_id": 3
// }

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
}

var events []Event

func main() {
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	// Middleware для логирования
	http.HandleFunc("/", logRequestHandler)

	// Запуск HTTP-сервера на порту 3000
	port := 3000
	fmt.Printf("Сервер запущен на порту %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func logRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Логирование входящего запроса
	fmt.Printf("Входящий запрос: %s %s\n", r.Method, r.URL.Path)
	http.DefaultServeMux.ServeHTTP(w, r)
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Чтение параметров из JSON в теле запроса
	var eventData struct {
		UserID  int    `json:"user_id"`
		Title   string `json:"title"`
		DateStr string `json:"date"`
	}

	err := json.NewDecoder(r.Body).Decode(&eventData)
	if err != nil {
		http.Error(w, "Ошибка в формате JSON", http.StatusBadRequest)
		return
	}

	userID := eventData.UserID
	title := eventData.Title

	date, err := time.Parse("2006-01-02", eventData.DateStr)
	if err != nil {
		http.Error(w, "Невалидный параметр date (ожидается формат 'YYYY-MM-DD')", http.StatusBadRequest)
		return
	}

	// Создание события и добавление его в список
	event := Event{
		ID:     len(events) + 1,
		UserID: userID,
		Title:  title,
		Date:   date,
	}

	events = append(events, event)

	// Отправка успешного ответа
	response := map[string]string{"result": "Событие создано успешно"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Чтение параметров из JSON в теле запроса
	var eventData struct {
		EventID int    `json:"event_id"`
		Title   string `json:"title"`
		DateStr string `json:"date"`
	}

	err := json.NewDecoder(r.Body).Decode(&eventData)
	if err != nil {
		http.Error(w, "Ошибка в формате JSON", http.StatusBadRequest)
		return
	}

	eventID := eventData.EventID
	title := eventData.Title

	date, err := time.Parse("2006-01-02", eventData.DateStr)
	if err != nil {
		http.Error(w, "Невалидный параметр date (ожидается формат 'YYYY-MM-DD')", http.StatusBadRequest)
		return
	}

	// Поиск события по ID и обновление его данных
	var updated bool
	for i, event := range events {
		if event.ID == eventID {
			events[i].Title = title
			events[i].Date = date
			updated = true
			break
		}
	}

	if !updated {
		http.Error(w, "Событие не найдено", http.StatusNotFound)
		return
	}

	// Отправка успешного ответа
	response := map[string]string{"result": "Событие обновлено успешно"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Чтение параметров из JSON в теле запроса
	var eventData struct {
		EventID int `json:"event_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&eventData)
	if err != nil {
		http.Error(w, "Ошибка в формате JSON", http.StatusBadRequest)
		return
	}

	eventID := eventData.EventID

	// Поиск события по ID и его удаление
	var found bool
	for i, event := range events {
		if event.ID == eventID {
			// Удаление события из списка
			events = append(events[:i], events[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Событие не найдено", http.StatusNotFound)
		return
	}

	// Отправка успешного ответа
	response := map[string]string{"result": "Событие удалено успешно"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Чтение параметров из query string
	dateStr := r.URL.Query().Get("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Невалидный параметр date (ожидается формат 'YYYY-MM-DD')", http.StatusBadRequest)
		return
	}

	// Фильтрация событий по дате и пользователю (если указан)
	userIDStr := r.URL.Query().Get("user_id")
	var userID int
	if userIDStr != "" {
		userID, err = strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Невалидный параметр user_id", http.StatusBadRequest)
			return
		}
	}

	var filteredEvents []Event
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			if userID == 0 || event.UserID == userID {
				filteredEvents = append(filteredEvents, event)
			}
		}
	}

	// Отправка ответа с событиями
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredEvents)
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Чтение параметров из query string
	dateStr := r.URL.Query().Get("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Невалидный параметр date (ожидается формат 'YYYY-MM-DD')", http.StatusBadRequest)
		return
	}

	// Определение начала и конца недели для заданной даты
	weekStart := date.AddDate(0, 0, -int(date.Weekday()))
	weekEnd := weekStart.AddDate(0, 0, 6)

	// Фильтрация событий по диапазону дат и пользователю (если указан)
	userIDStr := r.URL.Query().Get("user_id")
	var userID int
	if userIDStr != "" {
		userID, err = strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Невалидный параметр user_id", http.StatusBadRequest)
			return
		}
	}

	var filteredEvents []Event
	for _, event := range events {
		if event.Date.After(weekStart) && event.Date.Before(weekEnd) {
			if userID == 0 || event.UserID == userID {
				filteredEvents = append(filteredEvents, event)
			}
		}
	}

	// Отправка ответа с событиями
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredEvents)
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	// Чтение параметров из query string
	dateStr := r.URL.Query().Get("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Невалидный параметр date (ожидается формат 'YYYY-MM-DD')", http.StatusBadRequest)
		return
	}

	// Определение начала и конца месяца для заданной даты
	monthStart := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, -1)

	// Фильтрация событий по диапазону дат и пользователю (если указан)
	userIDStr := r.URL.Query().Get("user_id")
	var userID int
	if userIDStr != "" {
		userID, err = strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Невалидный параметр user_id", http.StatusBadRequest)
			return
		}
	}

	var filteredEvents []Event
	for _, event := range events {
		if event.Date.After(monthStart) && event.Date.Before(monthEnd) {
			if userID == 0 || event.UserID == userID {
				filteredEvents = append(filteredEvents, event)
			}
		}
	}

	// Отправка ответа с событиями
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredEvents)
}
