package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string ` json:"password"`
}
