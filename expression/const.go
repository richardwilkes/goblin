// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Const defines a constant.
type Const struct {
	Value reflect.Value
	ast.PosImpl
}

func (expr *Const) String() string {
	switch expr.Value {
	case ast.TrueValue:
		return "true"
	case ast.FalseValue:
		return "false"
	default:
		return "nil"
	}
}

// Invoke the expression and return a result.
func (expr *Const) Invoke(_ ast.Scope) (reflect.Value, error) {
	return expr.Value, nil
}

// Assign a value to the expression and return it.
func (expr *Const) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
