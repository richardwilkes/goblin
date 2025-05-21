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

// Make defines a make expression.
type Make struct {
	Type string
	ast.PosImpl
}

func (expr *Make) String() string {
	return fmt.Sprintf("make(%s)", expr.Type)
}

// Invoke the expression and return a result.
func (expr *Make) Invoke(scope ast.Scope) (reflect.Value, error) {
	rt, err := scope.Type(expr.Type)
	if err != nil {
		return ast.NilValue, ast.NewError(expr, err)
	}
	if rt.Kind() == reflect.Map {
		return reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt), nil
	}
	return reflect.Zero(rt), nil
}

// Assign a value to the expression and return it.
func (expr *Make) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
