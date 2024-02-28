package constants

import "errors"

type PqErrorCode string

const (
	DuplicateKeyViolationCode PqErrorCode = "23505"
	ForeignKeyViolationCode   PqErrorCode = "23503"
	UniqueViolationCode       PqErrorCode = "23514"
)

type PqErrors error

var (
	ErrorNotAffected PqErrors = errors.New("not affected")
	ErrorDuplicate   PqErrors = errors.New("duplicate entry")
	ErrForegignKey   PqErrors = errors.New("foreign key violation")
)
