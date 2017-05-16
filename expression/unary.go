package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// Unary defines a unary expression, e.g.: -1, ^1, ~1.
type Unary struct {
	goblin.PosImpl
	Operator string
	Expr     goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Unary) Invoke(env *goblin.Env) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(env)
	if err != nil {
		return v, goblin.NewError(expr, err)
	}
	switch expr.Operator {
	case "-":
		if v.Kind() == reflect.Float64 {
			return reflect.ValueOf(-v.Float()), nil
		}
		return reflect.ValueOf(-v.Int()), nil
	case "^":
		return reflect.ValueOf(^util.ToInt64(v)), nil
	case "!":
		return reflect.ValueOf(!util.ToBool(v)), nil
	default:
		return goblin.NilValue, goblin.NewStringError(expr, fmt.Sprintf("Unknown operator '%s'", expr.Operator))
	}
}

// Assign a value to the expression and return it.
func (expr *Unary) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
