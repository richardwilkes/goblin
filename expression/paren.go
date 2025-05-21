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

// Paren defines a parent block expression.
type Paren struct {
	SubExpr ast.Expr
	ast.PosImpl
}

func (expr *Paren) String() string {
	return fmt.Sprintf("(%v)", expr.SubExpr)
}

// Invoke the expression and return a result.
func (expr *Paren) Invoke(scope ast.Scope) (reflect.Value, error) {
	v, err := expr.SubExpr.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	return v, nil
}

// Assign a value to the expression and return it.
func (expr *Paren) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
