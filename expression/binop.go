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
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// BinOp defines a binary operator expression.
type BinOp struct {
	Left     ast.Expr
	Right    ast.Expr
	Operator string
	ast.PosImpl
}

func (expr *BinOp) String() string {
	return fmt.Sprintf("%v %s %v", expr.Left, expr.Operator, expr.Right)
}

// Invoke the expression and return a result.
func (expr *BinOp) Invoke(scope ast.Scope) (reflect.Value, error) {
	right := ast.NilValue
	left, err := expr.Left.Invoke(scope)
	if err != nil {
		return left, ast.NewError(expr, err)
	}
	if left.Kind() == reflect.Interface {
		left = left.Elem()
	}
	if expr.Right != nil {
		right, err = expr.Right.Invoke(scope)
		if err != nil {
			return right, ast.NewError(expr, err)
		}
		if right.Kind() == reflect.Interface {
			right = right.Elem()
		}
	}
	switch expr.Operator {
	case "+":
		if left.Kind() == reflect.String || right.Kind() == reflect.String {
			return reflect.ValueOf(util.ToString(left) + util.ToString(right)), nil
		}
		if (left.Kind() == reflect.Array || left.Kind() == reflect.Slice) && (right.Kind() != reflect.Array && right.Kind() != reflect.Slice) {
			return reflect.Append(left, right), nil
		}
		if (left.Kind() == reflect.Array || left.Kind() == reflect.Slice) && (right.Kind() == reflect.Array || right.Kind() == reflect.Slice) {
			return reflect.AppendSlice(left, right), nil
		}
		if left.Kind() == reflect.Float64 || right.Kind() == reflect.Float64 {
			return reflect.ValueOf(util.ToFloat64(left) + util.ToFloat64(right)), nil
		}
		return reflect.ValueOf(util.ToInt64(left) + util.ToInt64(right)), nil
	case "-":
		if left.Kind() == reflect.Float64 || right.Kind() == reflect.Float64 {
			return reflect.ValueOf(util.ToFloat64(left) - util.ToFloat64(right)), nil
		}
		return reflect.ValueOf(util.ToInt64(left) - util.ToInt64(right)), nil
	case "*":
		if left.Kind() == reflect.String && (right.Kind() == reflect.Int || right.Kind() == reflect.Int32 || right.Kind() == reflect.Int64) {
			return reflect.ValueOf(strings.Repeat(util.ToString(left), int(util.ToInt64(right)))), nil
		}
		if left.Kind() == reflect.Float64 || right.Kind() == reflect.Float64 {
			return reflect.ValueOf(util.ToFloat64(left) * util.ToFloat64(right)), nil
		}
		return reflect.ValueOf(util.ToInt64(left) * util.ToInt64(right)), nil
	case "/":
		return reflect.ValueOf(util.ToFloat64(left) / util.ToFloat64(right)), nil
	case "%":
		return reflect.ValueOf(util.ToInt64(left) % util.ToInt64(right)), nil
	case "==":
		return reflect.ValueOf(util.Equal(left, right)), nil
	case "!=":
		return reflect.ValueOf(!util.Equal(left, right)), nil
	case ">":
		return reflect.ValueOf(util.ToFloat64(left) > util.ToFloat64(right)), nil
	case ">=":
		return reflect.ValueOf(util.ToFloat64(left) >= util.ToFloat64(right)), nil
	case "<":
		return reflect.ValueOf(util.ToFloat64(left) < util.ToFloat64(right)), nil
	case "<=":
		return reflect.ValueOf(util.ToFloat64(left) <= util.ToFloat64(right)), nil
	case "|":
		return reflect.ValueOf(util.ToInt64(left) | util.ToInt64(right)), nil
	case "||":
		if util.ToBool(left) {
			return left, nil
		}
		return right, nil
	case "&":
		return reflect.ValueOf(util.ToInt64(left) & util.ToInt64(right)), nil
	case "&&":
		if util.ToBool(left) {
			return right, nil
		}
		return left, nil
	case "**":
		if left.Kind() == reflect.Float64 {
			return reflect.ValueOf(math.Pow(util.ToFloat64(left), util.ToFloat64(right))), nil
		}
		return reflect.ValueOf(int64(math.Pow(util.ToFloat64(left), util.ToFloat64(right)))), nil
	case ">>":
		return reflect.ValueOf(util.ToInt64(left) >> uint64(util.ToInt64(right))), nil
	case "<<":
		return reflect.ValueOf(util.ToInt64(left) << uint64(util.ToInt64(right))), nil
	default:
		return ast.NilValue, ast.NewStringError(expr, "Unknown operator")
	}
}

// Assign a value to the expression and return it.
func (expr *BinOp) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
