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
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Slice defines an array slice expression.
type Slice struct {
	Value ast.Expr
	Begin ast.Expr
	End   ast.Expr
	ast.PosImpl
}

func (expr *Slice) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%v[", expr.Value)
	if expr.Begin != nil {
		fmt.Fprint(&buffer, expr.Begin)
	}
	buffer.WriteString(":")
	if expr.End != nil {
		fmt.Fprint(&buffer, expr.End)
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Slice) Invoke(scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Value.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	kind := v.Kind()
	if kind != reflect.String && kind != reflect.Array && kind != reflect.Slice {
		return v, ast.NewInvalidOperationError(expr)
	}
	begin, end, err := expr.extractIndexes(v, scope)
	if err != nil {
		return v, err
	}
	if kind == reflect.String {
		if begin > v.Len() || end > v.Len() {
			return ast.NilValue, ast.NewIndexOutOfRangeError(expr)
		}
		return reflect.ValueOf(v.String()[begin:end]), nil
	}
	if begin > v.Cap() || end > v.Cap() {
		return ast.NilValue, ast.NewIndexOutOfRangeError(expr)
	}
	return v.Slice(begin, end), nil
}

// Assign a value to the expression and return it.
func (expr *Slice) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Value.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	kind := v.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return v, ast.NewInvalidOperationError(expr)
	}
	begin, end, err := expr.extractIndexes(v, scope)
	if err != nil {
		return v, err
	}
	if begin > v.Cap() || end > v.Cap() {
		return v, ast.NewIndexOutOfRangeError(expr)
	}
	vv := v.Slice(begin, end)
	if !vv.CanSet() {
		return v, ast.NewCannotAssignError(expr)
	}
	vv.Set(rv)
	return rv, nil
}

func (expr *Slice) extractIndexes(v reflect.Value, scope ast.Scope) (begin, end int, err error) {
	if expr.Begin != nil {
		if begin, err = expr.extractIndex(expr.Begin, scope); err != nil {
			return 0, 0, err
		}
	}
	if expr.End != nil {
		if end, err = expr.extractIndex(expr.End, scope); err != nil {
			return 0, 0, err
		}
	} else {
		end = v.Len()
	}
	if begin < 0 || end < 0 {
		return 0, 0, ast.NewIndexOutOfRangeError(expr)
	}
	if begin > end {
		return 0, 0, ast.NewStringError(expr, "Beginning index must be less than or equal to ending index")
	}
	return begin, end, nil
}

func (expr *Slice) extractIndex(vex ast.Expr, scope ast.Scope) (int, error) {
	value, err := vex.Invoke(scope)
	if err != nil {
		return 0, ast.NewError(expr, err)
	}
	kind := value.Kind()
	if kind != reflect.Int && kind != reflect.Int64 {
		return 0, ast.NewIndexShouldBeIntError(expr)
	}
	return int(value.Int()), nil
}
