package utils

import (
	"errors"

	"github.com/lib/pq"
)

func GetPostgresError(err error) (*pq.Error, bool) {
	var pqErr *pq.Error

	if errors.As(err, &pqErr) {
		return pqErr, true
	}
	return nil, false
}
