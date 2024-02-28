package utils

import (
	"fmt"
	"net/http"

	"github.com/lib/pq"
	"github.com/umardev500/chat/constants"
)

type PqCustomError struct {
	Message  string
	SQLState pq.ErrorCode
}

// Error returns the error message
func (e PqCustomError) Error() string {
	return fmt.Sprintf("%s (Code: %s)", e.Message, e.SQLState)
}

func ParsePostgresError(err error) error {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		return fmt.Errorf(http.StatusText(http.StatusInternalServerError))
	}

	code := pqErr.Code

	switch code {
	case pq.ErrorCode(constants.DuplicateKeyViolationCode):
		return PqCustomError{
			Message:  "duplicate key violation please check your data",
			SQLState: code,
		}
	case pq.ErrorCode(constants.ForeignKeyViolationCode):
		return PqCustomError{
			Message:  "foreign key violation please check your data",
			SQLState: code,
		}
	}

	return nil
}

func CombinePqErr(err error, dest *string) {
	newErr, ok := err.(PqCustomError)
	if ok {
		if dest != nil {
			msg := fmt.Sprintf("%s. %s", *dest, newErr.Message)
			*dest = msg
		}
	}

}
