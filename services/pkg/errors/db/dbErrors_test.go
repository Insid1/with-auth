package db_test

import (
	"errors"
	"testing"

	"github.com/Insid1/with-auth/pkg/errors/db"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var ErrSimple = errors.New("some error")

func TestCheckIsDBError_1(t *testing.T) {
	t.Parallel()

	result := db.CheckIsDBError(ErrSimple)

	assert.Nil(t, result)
}

func TestCheckIsDBError_2(t *testing.T) {
	t.Parallel()

	pqErr := pq.Error{
		Severity:         "",
		Code:             "",
		Message:          "realPq error",
		Detail:           "",
		Hint:             "",
		Position:         "",
		InternalPosition: "",
		InternalQuery:    "",
		Where:            "",
		Schema:           "",
		Table:            "",
		Column:           "",
		DataTypeName:     "",
		Constraint:       "",
		File:             "",
		Line:             "",
		Routine:          "",
	}
	result := db.CheckIsDBError(&pqErr)

	assert.NotNil(t, result)
	assert.Equal(t, "realPq error", result.Message)
}
