package model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// Содержит значения, которые переданы в полях
func TestBuildUpdateString_1(t *testing.T) {

	var usr User

	usr.ID = "id"
	usr.Email = "test email"
	usr.Name = "test name"
	usr.PassHash = "test_pass-hash"

	r, _ := usr.BuildUpdateString()

	assert.True(t, strings.Contains(r, usr.Email))
	assert.True(t, strings.Contains(r, usr.Name))
	assert.True(t, strings.Contains(r, usr.PassHash))
}

// Не содержит значения, которые переданы в полях
func TestBuildUpdateString_2(t *testing.T) {

	var usr User

	usr.ID = "id"
	usr.CreatedAt = time.Now()
	usr.UpdatedAt = time.Now()

	r, _ := usr.BuildUpdateString()

	assert.False(t, strings.Contains(r, usr.ID))
	assert.False(t, strings.Contains(r, usr.CreatedAt.GoString()))
	assert.False(t, strings.Contains(r, usr.UpdatedAt.GoString()))
}

// Возвращает ошибку
func TestBuildUpdateString_3(t *testing.T) {

	var usr User

	usr.Email = "test email"
	usr.Name = "test name"
	usr.PassHash = "test_pass-hash"

	r, err := usr.BuildUpdateString()

	assert.NotNil(t, err)
	assert.Equal(t, r, "")
}

// Проверки временной метки
func TestBuildUpdateString_4(t *testing.T) {
	var usr User

	usr.ID = "id"

	r, _ := usr.BuildUpdateString()

	assert.False(t, strings.Contains(r, "CURRENT_TIMESTAMP"))

	usr.Email = "test email"

	r, _ = usr.BuildUpdateString()

	assert.True(t, strings.Contains(r, "CURRENT_TIMESTAMP"))
}

// Проверка обновления хэша пароля по переданному паролю
func TestUpdatePassHash_1(t *testing.T) {
	var usr User

	assert.Empty(t, usr.PassHash)

	password := "hashed_password_here"

	err := usr.UpdatePassHash(password)

	assert.Nil(t, err)
	assert.NotEmpty(t, usr.PassHash)

	assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(usr.PassHash), []byte(password)))
}

// Проверка обновления хэша пароля при пустой строке пароля (Обновления не должно происходить)
func TestUpdatePassHash_2(t *testing.T) {
	var usr User

	assert.Empty(t, usr.PassHash)

	usr.PassHash = "some test data"

	password := ""

	err := usr.UpdatePassHash(password)

	assert.NotNil(t, err)
	assert.Equal(t, usr.PassHash, "some test data")
}
