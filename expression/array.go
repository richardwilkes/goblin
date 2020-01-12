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

// Array defines an array expression.
type Array struct {
	ast.PosImpl
	Exprs []ast.Expr
}

func (expr *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, v := range expr.Exprs {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprint(&buffer, v)
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Array) Invoke(scope ast.Scope) (reflect.Value, error) {
	a := make([]interface{}, len(expr.Exprs))
	for i, e := range expr.Exprs {
		arg, err := e.Invoke(scope)
		if err != nil {
			return arg, ast.NewError(e, err)
		}
		a[i] = arg.Interface()
	}
	return reflect.ValueOf(a), nil
}

// Assign a value to the expression and return it.
func (expr *Array) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
