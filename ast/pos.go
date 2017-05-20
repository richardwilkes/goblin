package ast

import (
	"fmt"
	"strconv"
)

// Position defines the position within source text.
type Position struct {
	Line   int
	Column int
}

func (p Position) String() string {
	if p.Column != 0 {
		return fmt.Sprintf("%d:%d", p.Line, p.Column)
	}
	return strconv.Itoa(p.Line)
}

// Pos defines the interface for getting and setting a position.
type Pos interface {
	// Position returns a position within source text.
	Position() Position
	// SetPosition sets a position within the source text.
	SetPosition(Position)
}

// PosImpl provides an implementation of the Pos interface that may be embedded.
type PosImpl struct {
	pos Position
}

// Position returns a position within source text.
func (x *PosImpl) Position() Position {
	return x.pos
}

// SetPosition sets a position within the source text.
func (x *PosImpl) SetPosition(pos Position) {
	x.pos = pos
}
