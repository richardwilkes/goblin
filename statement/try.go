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

// Try defines the try/catch/finally statement.
type Try struct {
	Var     string
	Try     []ast.Stmt
	Catch   []ast.Stmt
	Finally []ast.Stmt
	ast.PosImpl
}

func (stmt *Try) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("try {")
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Try {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n} catch ")
	if stmt.Var != "" {
		buffer.WriteString(stmt.Var)
		buffer.WriteString(" ")
	}
	buffer.WriteString("{")
	for _, one := range stmt.Catch {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	if len(stmt.Finally) > 0 {
		buffer.WriteString(" finally {")
		for _, one := range stmt.Finally {
			fmt.Fprintf(prefixer, "\n%v", one)
		}
		buffer.WriteString("\n}")
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Try) Execute(scope ast.Scope) (reflect.Value, error) {
	newScope := scope.NewScope()
	defer newScope.Destroy()
	_, err := newScope.Run(stmt.Try)
	if err != nil {
		// Catch
		catchScope := scope.NewScope()
		defer catchScope.Destroy()
		if stmt.Var != "" {
			catchScope.Define(stmt.Var, reflect.ValueOf(err))
		}
		_, e1 := catchScope.Run(stmt.Catch)
		if e1 != nil {
			err = ast.NewError(stmt.Catch[0], e1)
		} else {
			err = nil
		}
	}
	if len(stmt.Finally) > 0 {
		// Finally
		finallyScope := scope.NewScope()
		defer finallyScope.Destroy()
		_, e2 := finallyScope.Run(stmt.Finally)
		if e2 != nil {
			err = ast.NewError(stmt.Finally[0], e2)
		}
	}
	return ast.NilValue, ast.NewError(stmt, err)
}
