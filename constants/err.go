package constants

import "fmt"

type Error error

var (
	ErrNotAffected Error = fmt.Errorf("not affected")
)
