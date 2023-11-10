package handler

import (
	"11/pkg/storage"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type responseError struct {
	str  string
	code int
	err  error
}

// записываем дату из запроса или конкретную ошибку
func getData(r *http.Request) (*storage.RequestData, *responseError) {
	tmp := storage.RequestData{}
	if r.Method != http.MethodPost {
		return nil, &responseError{str: "404 Not Found", code: http.StatusNotFound, err: nil}
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, &responseError{str: "503 Server error", code: http.StatusServiceUnavailable, err: err}
	}
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return nil, &responseError{str: "400 Incorrect input data", code: http.StatusBadRequest, err: err}
	}
	tmp.DataTime, err = time.Parse("2006-01-02", tmp.DateJSON)
	if err != nil {
		{
			return nil, &responseError{str: "400 Incorrect input data", code: http.StatusBadRequest, err: err}
		}
	}
	return &tmp, nil
}

// добаляем ивент в хранилище
func createEvent(cash *storage.Cash) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, responseError := getData(r)
		if responseError != nil {
			log.Println(responseError.err)
			http.Error(w, responseError.str, responseError.code)
			return
		}
		if err := cash.Add(data); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("400 %s", err), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK) // возвращет статус 200
		w.Write([]byte(fmt.Sprintf("Успешно created мероприятие %s для пользователя в %s", data.User, data.DataTime)))
		log.Println("мероприятие добавленно")
	}
}

// обновляем данне
func updateEvent(cash *storage.Cash) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, responseError := getData(r)
		if responseError != nil {
			log.Println(responseError.err)
			http.Error(w, responseError.str, responseError.code)
			return
		}
		if err := cash.Update(data); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("400 %s", err), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Успешно update мероприятие %s для пользователя в %s", data.User, data.DataTime)))
	}
}

// удаляем данные
func deleteEvent(cash *storage.Cash) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, responseError := getData(r)
		if responseError != nil {
			log.Println(responseError.err)
			http.Error(w, responseError.str, responseError.code)
			return
		}
		if err := cash.Delete(data); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("400 %s", err), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Успешно delete мероприятие %s для пользователя в %s", data.User, data.DataTime)))
	}
}

// поулчаем ивенты в определнный день
func eventsForDay(cash *storage.Cash) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "404 Не найдена", http.StatusNotFound)
			return
		}
		query := r.URL.Query()                                   // поулчаем всю строку запроса
		user := query.Get("user_id")                             // получаем со строки запроса user id
		date, err := time.Parse("2006-01-02", query.Get("date")) // в каком виде должна быть дата
		if err != nil {
			http.Error(w, "400 Не коректные данные", http.StatusBadRequest)
			return
		}
		event, err := cash.FindDayEvent(user, date)
		if err != nil {
			http.Error(w, fmt.Sprintf("503 %s", err), http.StatusServiceUnavailable)
			return
		}
		b, err := json.Marshal(event) // упаковываем данные для передачи через json
		if err != nil {
			http.Error(w, "503 Ошибка сервера", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("День - %s\n", user)))
		w.Write(b) // выводи информацию на странице
	}
}

// поулчание мероприятий на неделе
func eventsForWeek(cash *storage.Cash) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "404 Не найдена", http.StatusNotFound)
			return
		}
		query := r.URL.Query()
		user := query.Get("user_id")
		date, err := time.Parse("2006-01-02", query.Get("date"))
		if err != nil {
			http.Error(w, "400 Не коректные данные", http.StatusBadRequest)
			return
		}
		event, err := cash.FindWeekEvent(user, date)
		if err != nil {
			http.Error(w, fmt.Sprintf("503 %s", err), http.StatusServiceUnavailable)
			return
		}
		b, err := json.Marshal(event)
		if err != nil {
			http.Error(w, "503 Ошибка сервера", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Неделя - %s\n", user)))
		w.Write(b)
	}
}

// мероприятия в месяц
func eventsForMonth(cash *storage.Cash) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "404 Не найдена", http.StatusNotFound)
			return
		}
		query := r.URL.Query()
		user := query.Get("user_id")
		date, err := time.Parse("2006-01-02", query.Get("date"))
		if err != nil {
			http.Error(w, "400 Не коректные данные", http.StatusBadRequest)
			return
		}
		event, err := cash.FindMonthEvent(user, date)
		if err != nil {
			http.Error(w, fmt.Sprintf("503 %s", err), http.StatusServiceUnavailable)
			return
		}
		b, err := json.Marshal(event)
		if err != nil {
			http.Error(w, "503 Ошибка сервера", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Месяц - %s\n", user)))
		w.Write(b)
	}
}
