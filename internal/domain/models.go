package domain

import "time"

type Records struct {
	ID          int    `json:"id"`
	Title       string `json:"title,max=55" example:"Title"`
	Description string `json:"description,max=255" example:"Description"`
}

type Users struct {
	ID       int    `json:"id"`
	Login    string `json:"login" binding:"required,min=1" example:"ivashka"`
	Username string `json:"username" binding:"required,min=1" example:"ivan"`
	Password string `json:"password" binding:"required,min=4" example:"1234"`
}

type Tokens struct {
	ID         int       `db:"id"`
	Token      string    `db:"token"`
	Expiration time.Time `db:"expiration"`
}
