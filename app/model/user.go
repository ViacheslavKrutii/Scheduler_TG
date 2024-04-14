package model

import "time"

type UserId string
type User struct {
	ID               UserId
	Name             string
	Location         time.Location
	AllEventsStorage map[EventID]*Event
}

func NewUser(id UserId, name string, location time.Location) User {
	Events := make(map[EventID]*Event, 0)
	return User{
		ID:               id,
		Name:             name,
		Location:         location,
		AllEventsStorage: Events,
	}
}

func (u *User) AddEvent(e *Event) {
	u.AllEventsStorage[e.ID] = e
}

func (u *User) DeleteEventByID(ID EventID) {
	delete(u.AllEventsStorage, ID)
}
