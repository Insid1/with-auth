package model

import (
	"fmt"
	"strings"
	"time"

	userErrors "github.com/Insid1/with-auth/pkg/errors/user"
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
		return "", userErrors.ErrUserIDNotProvided
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
		return userErrors.ErrEmptyPassword
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PassHash = string(passHash)

	return nil
}
