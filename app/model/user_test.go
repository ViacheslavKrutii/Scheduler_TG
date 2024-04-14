package model

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	location, _ := time.LoadLocation("America/Los_Angeles")
	got := NewUser("12", "Bob", *location)
	want := User{
		ID:               "12",
		Name:             "Bob",
		Location:         *location,
		AllEventsStorage: make(map[EventID]*Event),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestAddEvent(t *testing.T) {
	timeNow := time.Now()
	e := Event{
		ID:      "1",
		Title:   "wake up",
		Days:    []time.Weekday{time.Sunday},
		Time:    timeNow,
		Message: "message",
	}
	u := User{
		ID:               "12",
		Name:             "Bob",
		AllEventsStorage: make(map[EventID]*Event),
	}
	u.AddEvent(&e)

	// Check if the event was added to the user's event storage
	_, ok := u.AllEventsStorage[e.ID]
	if !ok {
		t.Errorf("Event was not added to user's event storage")
	}
}

func TestDeleteEvent(t *testing.T) {
	timeNow := time.Now()
	e := Event{
		ID:      "1",
		Title:   "wake up",
		Days:    []time.Weekday{time.Sunday},
		Time:    timeNow,
		Message: "message",
	}
	u := User{
		ID:               "12",
		Name:             "Bob",
		AllEventsStorage: map[EventID]*Event{"1": &e},
	}
	u.DeleteEventByID(e.ID)

	// Check if the event was deleted from the user's event storage
	_, ok := u.AllEventsStorage[e.ID]
	if ok {
		t.Errorf("Event was not deleted from user's event storage")
	}
}
