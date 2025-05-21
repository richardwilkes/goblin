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

// Variable defines a variable definition statement.
type Variable struct {
	Names []string
	Exprs []ast.Expr
	ast.PosImpl
}

func (stmt *Variable) String() string {
	var buffer bytes.Buffer
	for i, name := range stmt.Names {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(name)
	}
	buffer.WriteString(" = ")
	for i, one := range stmt.Exprs {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Variable) Execute(scope ast.Scope) (reflect.Value, error) {
	rvs := make([]reflect.Value, 0, len(stmt.Exprs))
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(scope)
		if err != nil {
			return rv, ast.NewError(expr, err)
		}
		rvs = append(rvs, rv)
	}
	result := make([]any, 0, len(rvs))
	for i, name := range stmt.Names {
		if i < len(rvs) {
			scope.Define(name, rvs[i])
			result = append(result, rvs[i].Interface())
		}
	}
	return reflect.ValueOf(result), nil
}
