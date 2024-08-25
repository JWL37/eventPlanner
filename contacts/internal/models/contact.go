package models

type Contact struct {
	ID    int `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}
