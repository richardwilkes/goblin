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

// CFor defines a C-style "for (;;)" statement.
type CFor struct {
	ast.PosImpl
	Expr1 ast.Expr
	Expr2 ast.Expr
	Expr3 ast.Expr
	Stmts []ast.Stmt
}

func (stmt *CFor) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "for %v; %v; %v {", stmt.Expr1, stmt.Expr2, stmt.Expr3)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *CFor) Execute(scope ast.Scope) (reflect.Value, error) {
	newScope := scope.NewScope()
	defer newScope.Destroy()
	_, err := stmt.Expr1.Invoke(newScope)
	if err != nil {
		return ast.NilValue, err
	}
	for {
		var fb reflect.Value
		if fb, err = stmt.Expr2.Invoke(newScope); err != nil {
			return ast.NilValue, err
		}
		if !util.ToBool(fb) {
			break
		}
		var rv reflect.Value
		if rv, err = newScope.Run(stmt.Stmts); err != nil {
			if err == ast.ErrBreak {
				break
			}
			if err == ast.ErrContinue {
				continue
			}
			if err == ast.ErrReturn {
				return rv, err
			}
			return rv, ast.NewError(stmt, err)
		}
		_, err = stmt.Expr3.Invoke(newScope)
		if err != nil {
			return ast.NilValue, err
		}
	}
	return ast.NilValue, nil
}
