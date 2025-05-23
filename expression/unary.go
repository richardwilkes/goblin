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

// Unary defines a unary expression, e.g.: -1, ^1, ~1.
type Unary struct {
	Expr     ast.Expr
	Operator string
	ast.PosImpl
}

func (expr *Unary) String() string {
	return fmt.Sprintf("%s%v", expr.Operator, expr.Expr)
}

// Invoke the expression and return a result.
func (expr *Unary) Invoke(scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	switch expr.Operator {
	case "-":
		if v.Kind() == reflect.Float64 {
			return reflect.ValueOf(-v.Float()), nil
		}
		return reflect.ValueOf(-v.Int()), nil
	case "^":
		return reflect.ValueOf(^util.ToInt64(v)), nil
	case "!":
		return reflect.ValueOf(!util.ToBool(v)), nil
	default:
		return ast.NilValue, ast.NewStringError(expr, fmt.Sprintf("Unknown operator '%s'", expr.Operator))
	}
}

// Assign a value to the expression and return it.
func (expr *Unary) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
