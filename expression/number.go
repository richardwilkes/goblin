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

// Number defines a number expression.
type Number struct {
	ast.PosImpl
	Value reflect.Value
	Err   error
}

func (expr *Number) String() string {
	switch expr.Value.Kind() {
	case reflect.Float64:
		return fmt.Sprint(expr.Value.Float())
	case reflect.Int64:
		return fmt.Sprint(expr.Value.Int())
	default:
		return "<nil>"
	}
}

// Invoke the expression and return a result.
func (expr *Number) Invoke(scope ast.Scope) (reflect.Value, error) {
	return expr.Value, expr.Err
}

// Assign a value to the expression and return it.
func (expr *Number) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
