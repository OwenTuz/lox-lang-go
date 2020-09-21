package exception

import (
	"fmt"
)

type Exception struct {
	// A simple error type with a description and the current line
	description string
	line        int
}

func New(line int, description string) *Exception {
	return &Exception{description, line}
}

func (exc *Exception) Error() string {
	return fmt.Sprintf("Error at line %v: %s", exc.line, exc.description)
}
