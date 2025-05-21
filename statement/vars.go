// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Variables defines a statement which defines multiple variables.
type Variables struct {
	Operator string
	Left     []ast.Expr
	Right    []ast.Expr
	ast.PosImpl
}

func (stmt *Variables) String() string {
	var buffer bytes.Buffer
	for i, one := range stmt.Left {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	buffer.WriteString(" ")
	buffer.WriteString(stmt.Operator)
	buffer.WriteString(" ")
	for i, one := range stmt.Right {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Variables) Execute(scope ast.Scope) (reflect.Value, error) {
	vs := make([]any, 0, len(stmt.Right))
	for _, right := range stmt.Right {
		rv, err := right.Invoke(scope)
		if err != nil {
			return rv, ast.NewError(right, err)
		}
		switch {
		case rv == ast.NilValue: //nolint:govet // Yes, we do want to compare against this specific value
			vs = append(vs, nil)
		case rv.IsValid() && rv.CanInterface():
			vs = append(vs, rv.Interface())
		default:
			vs = append(vs, nil)
		}
	}
	rvs := reflect.ValueOf(vs)
	if len(stmt.Left) > 1 && rvs.Len() == 1 {
		item := rvs.Index(0)
		if item.Kind() == reflect.Interface {
			item = item.Elem()
		}
		if item.Kind() == reflect.Slice {
			rvs = item
		}
	}
	for i, left := range stmt.Left {
		if i >= rvs.Len() {
			break
		}
		v := rvs.Index(i)
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		_, err := left.Assign(v, scope)
		if err != nil {
			return rvs, ast.NewError(left, err)
		}
	}
	if rvs.Len() == 1 {
		return rvs.Index(0), nil
	}
	return rvs, nil
}
