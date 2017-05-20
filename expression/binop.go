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
	ast.PosImpl
	Left     ast.Expr
	Operator string
	Right    ast.Expr
}

func (expr *BinOp) String() string {
	return fmt.Sprintf("%v %s %v", expr.Left, expr.Operator, expr.Right)
}

// Invoke the expression and return a result.
func (expr *BinOp) Invoke(scope ast.Scope) (reflect.Value, error) {
	left := ast.NilValue
	right := ast.NilValue
	var err error

	left, err = expr.Left.Invoke(scope)
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
func (expr *BinOp) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
