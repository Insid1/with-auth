package db

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrUnableToOpenConnection = errors.New("unable to Open DB Connection")
	ErrUnableToConnect        = errors.New("unable to connect to DB")
)

func CheckIsDBError(err error) *pq.Error {
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			return pqErr
		}
	}

	return nil
}

func CheckAlreadyExistError(err error) error {
	if pqErr := CheckIsDBError(err); pqErr != nil {
		if pqErr.Code.Name() == "unique_violation" {
			return pqErr
		}
	}

	return nil
}
