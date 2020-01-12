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
)

// Member defines a member reference expression.
type Member struct {
	ast.PosImpl
	Expr ast.Expr
	Name string
}

func (expr *Member) String() string {
	return fmt.Sprintf("%v.%s", expr.Expr, expr.Name)
}

// Invoke the expression and return a result.
func (expr *Member) Invoke(scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if v.IsValid() && v.CanInterface() {
		if vme, ok := v.Interface().(ast.Scope); ok {
			var m reflect.Value
			if m, err = vme.Get(expr.Name); err != nil || !m.IsValid() {
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
			}
			return m, nil
		}
	}

	m := v.MethodByName(expr.Name)
	if !m.IsValid() {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Struct:
			m = v.FieldByName(expr.Name)
			if !m.IsValid() {
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
			}
		case reflect.Map:
			m = v.MapIndex(reflect.ValueOf(expr.Name))
			if !m.IsValid() {
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
			}
		default:
			return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
		}
	}
	return m, nil
}

// Assign a value to the expression and return it.
func (expr *Member) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if !v.IsValid() {
		return ast.NilValue, ast.NewCannotAssignError(expr)
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		v = v.FieldByName(expr.Name)
		if !v.CanSet() {
			return ast.NilValue, ast.NewCannotAssignError(expr)
		}
		v.Set(rv)
	case reflect.Map:
		v.SetMapIndex(reflect.ValueOf(expr.Name), rv)
	default:
		if !v.CanSet() {
			return ast.NilValue, ast.NewCannotAssignError(expr)
		}
		v.Set(rv)
	}
	return v, nil
}
