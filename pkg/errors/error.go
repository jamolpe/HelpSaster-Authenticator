package errormng

import (
	"fmt"
	"time"
)

// ErrorMGN : base error definition
type ErrorMGN struct {
	When time.Time
	What string
}

func (e ErrorMGN) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

// NewError : creates a new error with the specified message
func NewError(msg string) error {
	return ErrorMGN{
		time.Now(),
		msg,
	}
}
