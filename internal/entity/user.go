package entity

type User struct {
	Id       string `json:"-"`
	Email    string `json:"name" binding:"required"`
	Name     string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Password string `json:"password" binding:"required"`
}
