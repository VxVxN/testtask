package event

import (
	"time"
)

type Event struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

func New(userID int, date time.Time) *Event {
	return &Event{
		UserID: userID,
		Date:   date,
	}
}

func (event *Event) IncludeInDay(year, yearDay int) bool {
	return event.Date.Year() == year && event.Date.YearDay() == yearDay
}

func (event *Event) IncludeInWeek(year, yearDay int, weekday time.Weekday) bool {
	return event.Date.Year() == year && event.Date.YearDay() >= yearDay-int(weekday) && event.Date.YearDay() <= yearDay
}

func (event *Event) IncludeInMonth(year int, month time.Month) bool {
	return event.Date.Year() == year && event.Date.Month() == month
}
