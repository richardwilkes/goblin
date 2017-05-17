package interpreter

import (
	"errors"
	"fmt"
)

// Error constants
var (
	ErrBreak     = errors.New("Unexpected break statement")
	ErrContinue  = errors.New("Unexpected continue statement")
	ErrReturn    = errors.New("Unexpected return statement")
	ErrInterrupt = errors.New("Execution interrupted")
	ErrBadSyntax = errors.New("Bad syntax")
)

// Error provides a convenient interface for handling runtime errors.
type Error struct {
	Message string
	Pos     Position
	Fatal   bool
}

// Error returns the error message.
func (e *Error) Error() string {
	return fmt.Sprintf("%s [%v]", e.Message, e.Pos)
}

// NewStringError makes an error with a message and position.
func NewStringError(pos Pos, err string) error {
	if pos == nil {
		return &Error{Message: err, Pos: Position{Line: 1, Column: 1}}
	}
	return &Error{Message: err, Pos: pos.Position()}
}

// NewErrorf makes an error with a formatted message and position.
func NewErrorf(pos Pos, format string, args ...interface{}) error {
	return &Error{Message: fmt.Sprintf(format, args...), Pos: pos.Position()}
}

// NewError makes an error out of an existing one.
func NewError(pos Pos, err error) error {
	if err == nil {
		return nil
	}
	if err == ErrBreak || err == ErrContinue || err == ErrReturn {
		return err
	}
	if pe, ok := err.(*Error); ok {
		return pe
	}
	return &Error{Message: err.Error(), Pos: pos.Position()}
}

// NewInvalidOperationError ...
func NewInvalidOperationError(pos Pos) error {
	return NewStringError(pos, "Invalid operation")
}

// NewNamedInvalidOperationError ...
func NewNamedInvalidOperationError(pos Pos, name string) error {
	return NewErrorf(pos, "Invalid operation '%s'", name)
}

// NewArrayIndexShouldBeIntError ...
func NewArrayIndexShouldBeIntError(pos Pos) error {
	return NewStringError(pos, "Array index should be int")
}

// NewCannotAssignError ...
func NewCannotAssignError(pos Pos) error {
	return NewStringError(pos, "Cannot assign")
}

// NewMapKeyShouldBeStringError ...
func NewMapKeyShouldBeStringError(pos Pos) error {
	return NewStringError(pos, "Map key should be string")
}

// NewCannotExecuteError ...
func NewCannotExecuteError(pos Pos) error {
	return NewStringError(pos, "Cannot execute")
}
