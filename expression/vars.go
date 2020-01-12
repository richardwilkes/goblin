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

// Vars defines an expression that defines multiple variables.
type Vars struct {
	ast.PosImpl
	Left     []ast.Expr
	Operator string
	Right    []ast.Expr
}

func (expr *Vars) String() string {
	var buffer bytes.Buffer
	for i, one := range expr.Left {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	buffer.WriteString(" ")
	buffer.WriteString(expr.Operator)
	buffer.WriteString(" ")
	for i, one := range expr.Right {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Vars) Invoke(scope ast.Scope) (reflect.Value, error) {
	rv := ast.NilValue
	var err error
	var vs []interface{}
	for _, Right := range expr.Right {
		rv, err = Right.Invoke(scope)
		if err != nil {
			return rv, ast.NewError(Right, err)
		}
		switch {
		case rv == ast.NilValue:
			vs = append(vs, nil)
		case rv.IsValid() && rv.CanInterface():
			vs = append(vs, rv.Interface())
		default:
			vs = append(vs, nil)
		}
	}
	rvs := reflect.ValueOf(vs)
	for i, Left := range expr.Left {
		if i >= rvs.Len() {
			break
		}
		v := rvs.Index(i)
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		_, err = Left.Assign(v, scope)
		if err != nil {
			return rvs, ast.NewError(Left, err)
		}
	}
	return rvs, nil
}

// Assign a value to the expression and return it.
func (expr *Vars) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
