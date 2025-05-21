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
)

// Call defines a calling expression.
type Call struct {
	Func     any
	Name     string
	SubExprs []ast.Expr
	VarArg   bool
	ast.PosImpl
}

func (expr *Call) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(expr.Name)
	buffer.WriteString("(")
	for i, arg := range expr.SubExprs {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprint(&buffer, arg)
	}
	if expr.VarArg {
		buffer.WriteString("...")
	}
	buffer.WriteString(")")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Call) Invoke(scope ast.Scope) (reflect.Value, error) {
	f := ast.NilValue
	if expr.Func != nil {
		var ok bool
		if f, ok = expr.Func.(reflect.Value); !ok {
			f = ast.NilValue
		}
	} else {
		var err error
		ff, err := scope.Get(expr.Name)
		if err != nil {
			return f, err
		}
		f = ff
	}
	_, isReflect := f.Interface().(ast.Func)

	args := make([]reflect.Value, 0, len(expr.SubExprs))
	for i, subExpr := range expr.SubExprs {
		arg, err := subExpr.Invoke(scope)
		if err != nil {
			return arg, ast.NewError(subExpr, err)
		}
		typ := f.Type()
		if i < typ.NumIn() {
			if !typ.IsVariadic() {
				it := typ.In(i)
				if arg.Kind().String() == "unsafe.Pointer" {
					arg = reflect.New(it).Elem()
				}
				switch {
				case arg.Kind() != it.Kind() && arg.IsValid() && arg.Type().ConvertibleTo(it):
					arg = arg.Convert(it)
				case arg.Kind() == reflect.Func:
					if _, isFunc := arg.Interface().(ast.Func); isFunc {
						rfunc := arg
						arg = reflect.MakeFunc(it, func(args []reflect.Value) []reflect.Value {
							for j := range args {
								args[j] = reflect.ValueOf(args[j])
							}
							rets := make([]reflect.Value, 0, it.NumOut())
							for _, v := range rfunc.Call(args)[:it.NumOut()] {
								if value, ok := v.Interface().(reflect.Value); ok {
									rets = append(rets, value)
								}
							}
							return rets
						})
					}
				case !arg.IsValid():
					arg = reflect.Zero(it)
				default:
				}
			}
		}
		if !arg.IsValid() {
			arg = ast.NilValue
		}

		if !isReflect {
			if expr.VarArg && i == len(expr.SubExprs)-1 {
				for j := 0; j < arg.Len(); j++ {
					args = append(args, arg.Index(j).Elem())
				}
			} else {
				args = append(args, arg)
			}
		} else {
			if arg.Kind() == reflect.Interface {
				arg = arg.Elem()
			}
			if expr.VarArg && i == len(expr.SubExprs)-1 {
				for j := 0; j < arg.Len(); j++ {
					args = append(args, reflect.ValueOf(arg.Index(j).Elem()))
				}
			} else {
				args = append(args, reflect.ValueOf(arg))
			}
		}
	}
	ret := ast.NilValue
	var err error
	fnc := func() {
		defer func() {
			if ex := recover(); ex != nil {
				if e, ok := ex.(error); ok {
					err = e
				} else {
					err = errors.New(fmt.Sprint(ex))
				}
			}
		}()
		if f.Kind() == reflect.Interface {
			f = f.Elem()
		}
		rets := f.Call(args)
		if isReflect {
			var ok bool
			ev := rets[1].Interface()
			if ev != nil {
				if err, ok = ev.(error); !ok {
					err = nil
				}
			}
			if ret, ok = rets[0].Interface().(reflect.Value); !ok {
				ret = ast.NilValue
			}
		} else {
			for i, subExpr := range expr.SubExprs {
				if ae, ok := subExpr.(*Addr); ok {
					if id, ok2 := ae.Expr.(*Ident); ok2 {
						_, err = id.Assign(args[i].Elem().Elem(), scope)
					}
				}
			}
			if f.Type().NumOut() == 1 {
				ret = rets[0]
			} else {
				var result []any
				for _, r := range rets {
					result = append(result, r.Interface())
				}
				ret = reflect.ValueOf(result)
			}
		}
	}
	fnc()
	if err != nil {
		return ret, ast.NewError(expr, err)
	}
	return ret, nil
}

// Assign a value to the expression and return it.
func (expr *Call) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
