package goblin

import (
	"fmt"
	"reflect"
)

// UnaryExpr defines a unary expression, e.g.: -1, ^1, ~1.
type UnaryExpr struct {
	PosImpl
	Operator string
	Expr     Expr
}

// Invoke the expression and return a result.
func (expr *UnaryExpr) Invoke(env *Env) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	switch expr.Operator {
	case "-":
		if v.Kind() == reflect.Float64 {
			return reflect.ValueOf(-v.Float()), nil
		}
		return reflect.ValueOf(-v.Int()), nil
	case "^":
		return reflect.ValueOf(^toInt64(v)), nil
	case "!":
		return reflect.ValueOf(!toBool(v)), nil
	default:
		return NilValue, NewStringError(expr, fmt.Sprintf("Unknown operator '%s'", expr.Operator))
	}
}

// Assign a value to the expression and return it.
func (expr *UnaryExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
