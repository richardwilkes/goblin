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
