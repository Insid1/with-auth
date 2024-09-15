package model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
