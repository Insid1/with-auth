package model

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Age       uint32
	PassHash  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BuildUpdateString() (string, error) {
	if u.ID == "" {
		return "", fmt.Errorf("ID is not provided")
	}

	var result string

	if u.Name != "" {
		result += fmt.Sprintf("username='%s',", u.Name)
	}
	if u.Email != "" {
		result += fmt.Sprintf("email='%s',", u.Email)
	}
	if u.PassHash != "" {
		result += fmt.Sprintf("password_hash='%s',", u.PassHash)
	}

	if result != "" {
		result += "updated_at=CURRENT_TIMESTAMP"
	}

	result, _ = strings.CutSuffix(result, ",")

	return result, nil

}
