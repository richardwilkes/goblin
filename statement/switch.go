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

// Switch defines a switch statement.
type Switch struct {
	ast.PosImpl
	Expr  ast.Expr
	Cases []ast.Stmt
}

func (stmt *Switch) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "switch %v {", stmt.Expr)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Cases {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *Switch) Execute(scope ast.Scope) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(stmt, err)
	}
	done := false
	var defaultStmt *Default
	for _, ss := range stmt.Cases {
		if ssd, ok := ss.(*Default); ok {
			defaultStmt = ssd
			continue
		}
		caseStmt, ok := ss.(*Case)
		if !ok {
			return ast.NilValue, ast.NewError(stmt, ast.ErrBadSyntax)
		}
		cv, lerr := caseStmt.Expr.Invoke(scope)
		if lerr != nil {
			return rv, ast.NewError(stmt, lerr)
		}
		if !util.Equal(rv, cv) {
			continue
		}
		rv, err = scope.Run(caseStmt.Stmts)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
		done = true
		break
	}
	if !done && defaultStmt != nil {
		rv, err = scope.Run(defaultStmt.Stmts)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
	}
	return rv, nil
}
