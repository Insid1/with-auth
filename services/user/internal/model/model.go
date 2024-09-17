package model

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	Username  string
	Email     string
	PassHash  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BuildUpdateString() (string, error) {
	if u.ID == "" {
		return "", fmt.Errorf("ID is not provided")
	}

	var result string

	if u.Username != "" {
		result += fmt.Sprintf("username='%s',", u.Username)
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

func (u *User) UpdatePassHash(password string) error {
	if password == "" {
		return fmt.Errorf(`empty password provided`)
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PassHash = string(passHash)

	return nil
}
