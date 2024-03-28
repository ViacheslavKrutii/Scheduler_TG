package model

import (
	"reflect"
	"testing"
	"time"
)

func TestNewEvent(t *testing.T) {
	timeNow := time.Now()
	got := NewEvent("1", "wake up", "message", []time.Weekday{1}, timeNow)
	want := Event{
		ID:      "1",
		Title:   "wake up",
		Days:    []time.Weekday{1},
		Time:    timeNow,
		Message: "message",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}

}
