// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ast

import (
	"errors"
	"fmt"
)

// Error constants
var (
	ErrBreak     = errors.New("unexpected break statement")
	ErrContinue  = errors.New("unexpected continue statement")
	ErrReturn    = errors.New("unexpected return statement")
	ErrInterrupt = errors.New("execution interrupted")
	ErrBadSyntax = errors.New("bad syntax")
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

// NewInvalidOperationError ....
func NewInvalidOperationError(pos Pos) error {
	return NewStringError(pos, "Invalid operation")
}

// NewNamedInvalidOperationError ....
func NewNamedInvalidOperationError(pos Pos, name string) error {
	return NewErrorf(pos, "Invalid operation '%s'", name)
}

// NewIndexShouldBeIntError ....
func NewIndexShouldBeIntError(pos Pos) error {
	return NewStringError(pos, "Index should be int")
}

// NewIndexOutOfRangeError ....
func NewIndexOutOfRangeError(pos Pos) error {
	return NewStringError(pos, "Index out of range")
}

// NewCannotAssignError ....
func NewCannotAssignError(pos Pos) error {
	return NewStringError(pos, "Cannot assign")
}

// NewMapKeyShouldBeStringError ....
func NewMapKeyShouldBeStringError(pos Pos) error {
	return NewStringError(pos, "Map key should be string")
}

// NewCannotExecuteError ....
func NewCannotExecuteError(pos Pos) error {
	return NewStringError(pos, "Cannot execute")
}
