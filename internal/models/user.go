package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"` // может или не может создавать мероприятия
}
