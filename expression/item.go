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

// Item defines an expression that refers to a map or array item.
type Item struct {
	Value ast.Expr
	Index ast.Expr
	ast.PosImpl
}

func (expr *Item) String() string {
	return fmt.Sprintf("%v[%v]", expr.Value, expr.Index)
}

// Invoke the expression and return a result.
func (expr *Item) Invoke(scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Value.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	i, err := expr.Index.Invoke(scope)
	if err != nil {
		return i, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return ast.NilValue, ast.NewIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return ast.NilValue, nil
		}
		return v.Index(ii), nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return ast.NilValue, ast.NewMapKeyShouldBeStringError(expr)
		}
		return v.MapIndex(i), nil
	}
	if v.Kind() == reflect.String {
		rs := []rune(v.String())
		ii := int(i.Int())
		if ii < 0 || ii >= len(rs) {
			return ast.NilValue, nil
		}
		return reflect.ValueOf(rs[ii]), nil
	}
	return v, ast.NewInvalidOperationError(expr)
}

// Assign a value to the expression and return it.
func (expr *Item) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Value.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	i, err := expr.Index.Invoke(scope)
	if err != nil {
		return i, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return ast.NilValue, ast.NewIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return ast.NilValue, ast.NewCannotAssignError(expr)
		}
		vv := v.Index(ii)
		if !vv.CanSet() {
			return ast.NilValue, ast.NewCannotAssignError(expr)
		}
		vv.Set(rv)
		return rv, nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return ast.NilValue, ast.NewMapKeyShouldBeStringError(expr)
		}
		v.SetMapIndex(i, rv)
		return rv, nil
	}
	return v, ast.NewInvalidOperationError(expr)
}
