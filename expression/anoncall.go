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

// AnonCall defines an anonymous calling expression, e.g. func(){}().
type AnonCall struct {
	Expr     ast.Expr
	SubExprs []ast.Expr
	VarArg   bool
	ast.PosImpl
}

func (expr *AnonCall) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%v(", expr.Expr)
	for i, arg := range expr.SubExprs {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprint(&buffer, arg)
	}
	if expr.VarArg {
		buffer.WriteString("...")
	}
	buffer.WriteString(")")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *AnonCall) Invoke(scope ast.Scope) (reflect.Value, error) {
	f, err := expr.Expr.Invoke(scope)
	if err != nil {
		return f, ast.NewError(expr, err)
	}
	if f.Kind() == reflect.Interface {
		f = f.Elem()
	}
	if f.Kind() != reflect.Func {
		return f, ast.NewStringError(expr, "Unknown function")
	}
	call := &Call{Func: f, SubExprs: expr.SubExprs, VarArg: expr.VarArg}
	return call.Invoke(scope)
}

// Assign a value to the expression and return it.
func (expr *AnonCall) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
