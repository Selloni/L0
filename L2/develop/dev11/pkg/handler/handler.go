package handler

import (
	"11/config"
	"11/pkg/storage"
	"log"
	"net/http"
)

// Route маршруты для веб страницы вызывается в майн
func Route(cash *storage.Cash) error {
	conf := config.GetConfig()
	http.HandleFunc("/create_event", middleware(createEvent(cash)))
	http.HandleFunc("/update_event", middleware(updateEvent(cash)))
	http.HandleFunc("/delete_event", middleware(deleteEvent(cash)))
	http.HandleFunc("/events_for_day", middleware(eventsForDay(cash)))
	http.HandleFunc("/events_for_week", middleware(eventsForWeek(cash)))
	http.HandleFunc("/events_for_month", middleware(eventsForMonth(cash)))

	log.Println("инициализировали хандлер")
	err := http.ListenAndServe(":"+conf.Port, nil)
	if err != nil {
		return err
	}
	return nil
}
