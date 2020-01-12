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
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// TernaryOp defines a ternary operator expression.
type TernaryOp struct {
	ast.PosImpl
	Expr  ast.Expr
	Left  ast.Expr
	Right ast.Expr
}

func (expr *TernaryOp) String() string {
	return fmt.Sprintf("%v ? %v : %v", expr.Expr, expr.Left, expr.Right)
}

// Invoke the expression and return a result.
func (expr *TernaryOp) Invoke(scope ast.Scope) (reflect.Value, error) {
	rv, err := expr.Expr.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(expr, err)
	}
	var choice ast.Expr
	if util.ToBool(rv) {
		choice = expr.Left
	} else {
		choice = expr.Right
	}
	rv, err = choice.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(choice, err)
	}
	return rv, nil
}

// Assign a value to the expression and return it.
func (expr *TernaryOp) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
