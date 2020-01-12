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

// Loop defines a loop statement.
type Loop struct {
	ast.PosImpl
	Expr  ast.Expr
	Stmts []ast.Stmt
}

func (stmt *Loop) String() string {
	var buffer bytes.Buffer
	if stmt.Expr != nil {
		fmt.Fprintf(&buffer, "for %v {", stmt.Expr)
	} else {
		buffer.WriteString("for {")
	}
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *Loop) Execute(scope ast.Scope) (reflect.Value, error) {
	newScope := scope.NewScope()
	defer newScope.Destroy()
	for {
		if stmt.Expr != nil {
			ev, ee := stmt.Expr.Invoke(newScope)
			if ee != nil {
				return ev, ee
			}
			if !util.ToBool(ev) {
				break
			}
		}

		rv, err := newScope.Run(stmt.Stmts)
		if err != nil {
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
	}
	return ast.NilValue, nil
}
