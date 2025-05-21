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
	"github.com/richardwilkes/goblin/util"
)

// Return defines the return statement.
type Return struct {
	Exprs []ast.Expr
	ast.PosImpl
}

func (stmt *Return) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("return")
	for i, one := range stmt.Exprs {
		if i != 0 {
			buffer.WriteString(",")
		}
		fmt.Fprintf(&buffer, " %v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Return) Execute(scope ast.Scope) (reflect.Value, error) {
	var rvs []any
	switch len(stmt.Exprs) {
	case 0:
		return ast.NilValue, ast.ErrReturn
	case 1:
		rv, err := stmt.Exprs[0].Invoke(scope)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
		return rv, ast.ErrReturn
	}
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(scope)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
		switch {
		case util.IsNil(rv):
			rvs = append(rvs, nil)
		case rv.IsValid():
			rvs = append(rvs, rv.Interface())
		default:
			rvs = append(rvs, nil)
		}
	}
	return reflect.ValueOf(rvs), ast.ErrReturn
}
