package model

import "time"

type EventID string
type Event struct {
	ID      EventID        // ID
	Title   string         // Title
	Days    []time.Weekday // Choosen days
	Time    time.Time      // In which time must be alert
	Message string         // Message
}

func NewEvent(id EventID, title, message string, days []time.Weekday, time time.Time) Event {
	return Event{
		ID:      id,
		Title:   title,
		Days:    days,
		Time:    time,
		Message: message,
	}
}

