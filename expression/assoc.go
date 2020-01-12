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

// Assoc defines an operator association expression.
type Assoc struct {
	ast.PosImpl
	Left     ast.Expr
	Operator string
	Right    ast.Expr
}

func (expr *Assoc) String() string {
	switch expr.Operator {
	case "++", "--":
		return fmt.Sprintf("%v%s", expr.Left, expr.Operator)
	default:
		return fmt.Sprintf("%v %s %v", expr.Left, expr.Operator, expr.Right)
	}
}

// Invoke the expression and return a result.
func (expr *Assoc) Invoke(scope ast.Scope) (reflect.Value, error) {
	switch expr.Operator {
	case "++":
		if left, ok := expr.Left.(*Ident); ok {
			return expr.applyDelta(left, 1, scope)
		}
	case "--":
		if left, ok := expr.Left.(*Ident); ok {
			return expr.applyDelta(left, -1, scope)
		}
	}

	binop := &BinOp{Left: expr.Left, Operator: expr.Operator[0:1], Right: expr.Right}
	v, err := binop.Invoke(scope)
	if err != nil {
		return v, err
	}

	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return expr.Left.Assign(v, scope)
}

func (expr *Assoc) applyDelta(ident *Ident, delta int, scope ast.Scope) (reflect.Value, error) {
	v, err := scope.Get(ident.Lit)
	if err != nil {
		return v, err
	}
	if v.Kind() == reflect.Float64 {
		v = reflect.ValueOf(util.ToFloat64(v) + float64(delta))
	} else {
		v = reflect.ValueOf(util.ToInt64(v) + int64(delta))
	}
	if scope.Set(ident.Lit, v) != nil {
		scope.Define(ident.Lit, v)
	}
	return v, nil
}

// Assign a value to the expression and return it.
func (expr *Assoc) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
