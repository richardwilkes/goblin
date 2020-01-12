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

// MakeArray defines a make array expression.
type MakeArray struct {
	ast.PosImpl
	Type    string
	LenExpr ast.Expr
	CapExpr ast.Expr
}

func (expr *MakeArray) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("make([]")
	buffer.WriteString(expr.Type)
	if expr.LenExpr != nil {
		fmt.Fprintf(&buffer, ",%v", expr.LenExpr)
	}
	if expr.CapExpr != nil {
		fmt.Fprintf(&buffer, ",%v", expr.CapExpr)
	}
	buffer.WriteString(")")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *MakeArray) Invoke(scope ast.Scope) (reflect.Value, error) {
	typ, err := scope.Type(expr.Type)
	if err != nil {
		return ast.NilValue, err
	}
	var alen int
	if expr.LenExpr != nil {
		rv, lerr := expr.LenExpr.Invoke(scope)
		if lerr != nil {
			return ast.NilValue, lerr
		}
		alen = int(util.ToInt64(rv))
	}
	var acap int
	if expr.CapExpr != nil {
		rv, lerr := expr.CapExpr.Invoke(scope)
		if lerr != nil {
			return ast.NilValue, lerr
		}
		acap = int(util.ToInt64(rv))
	} else {
		acap = alen
	}
	return func() (reflect.Value, error) {
		defer func() {
			if ex := recover(); ex != nil {
				if e, ok := ex.(error); ok {
					err = e
				} else {
					err = errors.New(fmt.Sprint(ex))
				}
			}
		}()
		return reflect.MakeSlice(reflect.SliceOf(typ), alen, acap), nil
	}()
}

// Assign a value to the expression and return it.
func (expr *MakeArray) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
