package model_test

import (
	"strings"
	"testing"
	"time"

	"github.com/Insid1/with-auth/user/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// Содержит значения, которые переданы в полях.
func TestBuildUpdateString_1(t *testing.T) {
	t.Parallel()

	var usr model.User

	usr.ID = "id"
	usr.Email = "test email1"
	usr.Username = "test Username"
	usr.PassHash = "test_pass-hash"

	r, _ := usr.BuildUpdateString()

	assert.True(t, strings.Contains(r, usr.Email))
	assert.True(t, strings.Contains(r, usr.Username))
	assert.True(t, strings.Contains(r, usr.PassHash))
}

// Не содержит значения, которые переданы в полях.
func TestBuildUpdateString_2(t *testing.T) {
	t.Parallel()

	var usr model.User

	usr.ID = "id"
	usr.CreatedAt = time.Now()
	usr.UpdatedAt = time.Now()

	r, _ := usr.BuildUpdateString()

	assert.False(t, strings.Contains(r, usr.ID))
	assert.False(t, strings.Contains(r, usr.CreatedAt.GoString()))
	assert.False(t, strings.Contains(r, usr.UpdatedAt.GoString()))
}

// Возвращает ошибку.
func TestBuildUpdateString_3(t *testing.T) {
	t.Parallel()

	var usr model.User

	usr.Email = "test email"
	usr.Username = "test name"
	usr.PassHash = "test_pass-hash"

	res, err := usr.BuildUpdateString()

	require.Error(t, err)
	assert.Equal(t, "", res)
}

// Проверки временной метки.
func TestBuildUpdateString_4(t *testing.T) {
	t.Parallel()

	var usr model.User

	usr.ID = "id"

	res, _ := usr.BuildUpdateString()

	assert.False(t, strings.Contains(res, "CURRENT_TIMESTAMP"))

	usr.Email = "test email"

	res, _ = usr.BuildUpdateString()

	assert.True(t, strings.Contains(res, "CURRENT_TIMESTAMP"))
}

// Проверка обновления хэша пароля по переданному паролю.
func TestUpdatePassHash_1(t *testing.T) {
	t.Parallel()

	var usr model.User

	assert.Empty(t, usr.PassHash)

	password := "hashed_password_here"

	err := usr.UpdatePassHash(password)

	require.NoError(t, err)
	assert.NotEmpty(t, usr.PassHash)

	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(usr.PassHash), []byte(password)))
}

// Проверка обновления хэша пароля при пустой строке пароля (Обновления не должно происходить).
func TestUpdatePassHash_2(t *testing.T) {
	t.Parallel()

	var usr model.User

	assert.Empty(t, usr.PassHash)

	usr.PassHash = "some test data"

	password := ""

	err := usr.UpdatePassHash(password)

	require.Error(t, err)
	assert.Equal(t, "some test data", usr.PassHash)
}
