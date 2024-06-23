package model

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Age       uint32
	PassHash  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
