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

// Var defines an expression that defines a variable.
type Var struct {
	ast.PosImpl
	Left  ast.Expr
	Right ast.Expr
}

// Invoke the expression and return a result.
func (expr *Var) Invoke(scope ast.Scope) (reflect.Value, error) {
	rv, err := expr.Right.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(expr, err)
	}
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return expr.Left.Assign(rv, scope)
}

// Assign a value to the expression and return it.
func (expr *Var) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
