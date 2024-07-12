package models

import (
    "time"
)


type Event struct {
	ID        int    `json:"id"`
	NameEvent string `json:"name"`
	Shape     string `json:"shape"`
	Place     string `json:"place"`
	BeginTime time.Time `json:"begin_time"`
	Duration  string `json:"duration"`
	//listMembers []Contact `json:"members"`
}
