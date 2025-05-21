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
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Func defines a function expression.
type Func struct {
	Name   string
	Stmts  []ast.Stmt
	Args   []string
	VarArg bool
	ast.PosImpl
}

func (expr *Func) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("func")
	if expr.Name != "" {
		buffer.WriteString(" ")
		buffer.WriteString(expr.Name)
	}
	buffer.WriteString("(")
	for i, arg := range expr.Args {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(arg)
	}
	buffer.WriteString(") {")
	if len(expr.Stmts) > 0 {
		prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
		for _, stmt := range expr.Stmts {
			fmt.Fprintf(prefixer, "\n%v", stmt)
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Func) Invoke(scope ast.Scope) (reflect.Value, error) {
	f := reflect.ValueOf(func(fe *Func, scope ast.Scope) ast.Func {
		return func(args ...reflect.Value) (reflect.Value, error) {
			if !fe.VarArg {
				if len(args) != len(fe.Args) {
					return ast.NilValue, ast.NewStringError(fe, fmt.Sprintf("Expecting %d arguments, got %d", len(fe.Args), len(args)))
				}
			}
			newScope := scope.NewScope()
			if fe.VarArg {
				newScope.Define(fe.Args[0], reflect.ValueOf(args))
			} else {
				for i, arg := range fe.Args {
					newScope.Define(arg, args[i])
				}
			}
			rr, err := newScope.Run(fe.Stmts)
			if errors.Is(err, ast.ErrReturn) {
				err = nil
			}
			return rr, err
		}
	}(expr, scope))
	scope.Define(expr.Name, f)
	return f, nil
}

// Assign a value to the expression and return it.
func (expr *Func) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
