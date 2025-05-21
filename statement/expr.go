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
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Expression defines an expression statement.
type Expression struct {
	Expr ast.Expr
	ast.PosImpl
}

func (stmt *Expression) String() string {
	return fmt.Sprint(stmt.Expr)
}

// Execute the statement.
func (stmt *Expression) Execute(scope ast.Scope) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(stmt, err)
	}
	return rv, nil
}
