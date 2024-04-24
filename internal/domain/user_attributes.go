package domain

type Users struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required,min=1" example:"Ivan"`
	Username string `json:"username" binding:"required,min=1" example:"ivan"`
	Password string `json:"password" binding:"required,min=4" example:"1234"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,min=1" example:"Ivan"`
	Password string `json:"password" binding:"required,min=4" example:"1234"`
}
