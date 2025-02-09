package event

import (
	"time"
)

// Event todo
type Event struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

// New todo
func New(userID int, date time.Time) *Event {
	return &Event{
		UserID: userID,
		Date:   date,
	}
}

// IncludeInDay todo
func (event *Event) IncludeInDay(year, yearDay int) bool {
	return event.Date.Year() == year && event.Date.YearDay() == yearDay
}

// IncludeInWeek todo
func (event *Event) IncludeInWeek(year, yearDay int, weekday time.Weekday) bool {
	return event.Date.Year() == year && event.Date.YearDay() >= yearDay-int(weekday) && event.Date.YearDay() <= yearDay
}

// IncludeInMonth todo
func (event *Event) IncludeInMonth(year int, month time.Month) bool {
	return event.Date.Year() == year && event.Date.Month() == month
}
