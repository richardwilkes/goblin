package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Unary defines a unary expression, e.g.: -1, ^1, ~1.
type Unary struct {
	interpreter.PosImpl
	Operator string
	Expr     interpreter.Expr
}

func (expr *Unary) String() string {
	return fmt.Sprintf("%s%v", expr.Operator, expr.Expr)
}

// Invoke the expression and return a result.
func (expr *Unary) Invoke(env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
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
		return interpreter.NilValue, interpreter.NewStringError(expr, fmt.Sprintf("Unknown operator '%s'", expr.Operator))
	}
}

// Assign a value to the expression and return it.
func (expr *Unary) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
