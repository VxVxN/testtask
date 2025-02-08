package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/VxVxN/testtask/pkg/httphelper"
)

type Event struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

type EventResponse struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

// global storage for events
// todo use redis or something like that
var events = []Event{}

func main() {
	http.Handle("POST /event", logMiddleware(http.HandlerFunc(createEventHandler)))
	http.Handle("PUT /event", logMiddleware(http.HandlerFunc(updateEventHandler)))
	http.Handle("DELETE /event", logMiddleware(http.HandlerFunc(deleteEventHandler)))
	http.Handle("GET /events", logMiddleware(http.HandlerFunc(getEventsHandler)))

	fmt.Println("Server starting on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func parseEventParams(r *http.Request) (Event, error) {
	rawData, err := io.ReadAll(r.Body) // without r.FormValue() because it does not support DELETE requests
	if err != nil {
		return Event{}, err
	}
	values, err := url.ParseQuery(string(rawData))
	if err != nil {
		return Event{}, err
	}

	userID, err := strconv.Atoi(values.Get("user_id"))
	if err != nil {
		return Event{}, fmt.Errorf("invalid user_id")
	}

	date, err := time.Parse("2006-01-02", values.Get("date"))
	if err != nil {
		return Event{}, fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}

	return Event{UserID: userID, Date: date}, nil
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		endTime := time.Now()
		log.Printf("Completed request: %s %s, %s", r.Method, r.URL.Path, endTime.Sub(startTime))
	})
}

// todo create controller and move this
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventParams(r)
	if err != nil {
		httphelper.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	events = append(events, event)

	httphelper.SuccessResponse(w, "event created")
}

// todo create controller and move this
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventParams(r)
	if err != nil {
		httphelper.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	for i, e := range events {
		if e.UserID == event.UserID && e.Date == event.Date {
			events[i] = event // Update the event
			httphelper.SuccessResponse(w, "event updated")
			return
		}
	}

	httphelper.ErrorResponse(w, errors.New("event not found"), http.StatusBadRequest)
}

// todo create controller and move this
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventParams(r)
	if err != nil {
		httphelper.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	for i, e := range events {
		if e.UserID == event.UserID && e.Date == event.Date {
			events = append(events[:i], events[i+1:]...) // Delete the event
			httphelper.SuccessResponse(w, "event deleted")
			return
		}
	}

	httphelper.ErrorResponse(w, errors.New("event not found"), http.StatusServiceUnavailable)
}

func getEventsHandler(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")
	if period == "" {
		httphelper.ErrorResponse(w, errors.New("missing period parameter"), http.StatusBadRequest)
		return
	}

	periodEvents := []EventResponse{}
	now := time.Now()

	switch period {
	case "day":
		for _, e := range events {
			if e.Date.Year() == now.Year() && e.Date.YearDay() == now.YearDay() {
				periodEvents = append(periodEvents, EventResponse{
					UserID: e.UserID,
					Date:   e.Date.Format("2006-01-02"),
				})
			}
		}
	case "week":
		for _, e := range events {
			if e.Date.Year() == now.Year() && e.Date.YearDay() >= now.YearDay()-int(now.Weekday()) && e.Date.YearDay() <= now.YearDay() {
				periodEvents = append(periodEvents, EventResponse{
					UserID: e.UserID,
					Date:   e.Date.Format("2006-01-02"),
				})
			}
		}
	case "month":
		for _, e := range events {
			if e.Date.Year() == now.Year() && e.Date.Month() == now.Month() {
				periodEvents = append(periodEvents, EventResponse{
					UserID: e.UserID,
					Date:   e.Date.Format("2006-01-02"),
				})
			}
		}
	default:
		httphelper.ErrorResponse(w, errors.New("invalid period, expected: day, week, month"), http.StatusBadRequest)
		return
	}

	httphelper.SuccessResponse(w, periodEvents)
}
