package model

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Age       uint32
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Login struct {
	AppID    int32
	Email    string
	Password string
}

type Register struct {
	Email    string
	Password string
}
