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

// New defines a new instance expression.
type New struct {
	Type string
	ast.PosImpl
}

func (expr *New) String() string {
	return fmt.Sprintf("new(%s)", expr.Type)
}

// Invoke the expression and return a result.
func (expr *New) Invoke(scope ast.Scope) (reflect.Value, error) {
	rt, err := scope.Type(expr.Type)
	if err != nil {
		return ast.NilValue, ast.NewError(expr, err)
	}
	return reflect.New(rt), nil
}

// Assign a value to the expression and return it.
func (expr *New) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
