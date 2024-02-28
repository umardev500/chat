package constants

type ContextKey int

const (
	TransactionKey ContextKey = iota
)

// DeleteType is used for soft and hard delete
type DeleteType string

const (
	SoftDelete DeleteType = "soft"
	HardDelete DeleteType = "hard"
)
