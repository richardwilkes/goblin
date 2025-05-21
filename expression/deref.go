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

// Deref defines a dereferencing address expression.
type Deref struct {
	Expr ast.Expr
	ast.PosImpl
}

func (expr *Deref) String() string {
	return fmt.Sprintf("*%v", expr.Expr)
}

// Invoke the expression and return a result.
func (expr *Deref) Invoke(scope ast.Scope) (reflect.Value, error) {
	var v reflect.Value
	var err error
	switch ee := expr.Expr.(type) {
	case *Ident:
		if v, err = scope.Get(ee.Lit); err != nil {
			return v, err
		}
	case *Member:
		if v, err = ee.Expr.Invoke(scope); err != nil {
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
				if m, err = vme.Get(ee.Name); err != nil || !m.IsValid() {
					return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
				}
				return m, nil
			}
		}

		m := v.MethodByName(ee.Name)
		if !m.IsValid() {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			switch v.Kind() {
			case reflect.Struct:
				m = v.FieldByName(ee.Name)
				if !m.IsValid() {
					return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
				}
			case reflect.Map:
				m = v.MapIndex(reflect.ValueOf(ee.Name))
				if !m.IsValid() {
					return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
				}
			default:
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
			}
			v = m
		} else {
			v = m
		}
	default:
		return ast.NilValue, ast.NewStringError(expr, "Invalid operation for the value")
	}
	if v.Kind() != reflect.Ptr {
		return ast.NilValue, ast.NewStringError(expr, "Cannot deference for the value")
	}
	return v.Elem(), nil
}

// Assign a value to the expression and return it.
func (expr *Deref) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
