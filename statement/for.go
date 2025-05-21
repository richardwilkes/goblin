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
	"errors"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// For defines a for statement.
type For struct {
	Var   string
	Value ast.Expr
	Stmts []ast.Stmt
	ast.PosImpl
}

func (stmt *For) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "for %s in %v {", stmt.Var, stmt.Value)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *For) Execute(scope ast.Scope) (reflect.Value, error) {
	val, ee := stmt.Value.Invoke(scope)
	if ee != nil {
		return val, ee
	}
	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		newScope := scope.NewScope()
		defer newScope.Destroy()

		for i := 0; i < val.Len(); i++ {
			iv := val.Index(i)
			if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
				iv = iv.Elem()
			}
			newScope.Define(stmt.Var, iv)
			rv, err := newScope.Run(stmt.Stmts)
			if err != nil {
				if errors.Is(err, ast.ErrBreak) {
					break
				}
				if errors.Is(err, ast.ErrContinue) {
					continue
				}
				if errors.Is(err, ast.ErrReturn) {
					return rv, err
				}
				return rv, ast.NewError(stmt, err)
			}
		}
		return ast.NilValue, nil
	}
	return ast.NilValue, ast.NewStringError(stmt, "Invalid operation for non-array value")
}
