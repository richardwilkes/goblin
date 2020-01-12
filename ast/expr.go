// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ast

import "reflect"

// Expr defines the required methods of an expression.
type Expr interface {
	Pos
	// Invoke the expression and return a result.
	Invoke(scope Scope) (reflect.Value, error)
	// Assign a value to the expression and return it.
	Assign(rv reflect.Value, scope Scope) (reflect.Value, error)
}

// Type defines a type.
type Type struct {
	Name string
}
