package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Assoc defines an operator association expression.
type Assoc struct {
	interpreter.PosImpl
	Left     interpreter.Expr
	Operator string
	Right    interpreter.Expr
}

func (expr *Assoc) String() string {
	switch expr.Operator {
	case "++", "--":
		return fmt.Sprintf("%v%s", expr.Left, expr.Operator)
	default:
		return fmt.Sprintf("%v %s %v", expr.Left, expr.Operator, expr.Right)
	}
}

// Invoke the expression and return a result.
func (expr *Assoc) Invoke(env *interpreter.Env) (reflect.Value, error) {
	switch expr.Operator {
	case "++":
		if aLHS, ok := expr.Left.(*Ident); ok {
			v, err := env.Get(aLHS.Lit)
			if err != nil {
				return v, err
			}
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(util.ToFloat64(v) + 1.0)
			} else {
				v = reflect.ValueOf(util.ToInt64(v) + 1)
			}
			if env.Set(aLHS.Lit, v) != nil {
				env.Define(aLHS.Lit, v)
			}
			return v, nil
		}
	case "--":
		if aLHS, ok := expr.Left.(*Ident); ok {
			v, err := env.Get(aLHS.Lit)
			if err != nil {
				return v, err
			}
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(util.ToFloat64(v) - 1.0)
			} else {
				v = reflect.ValueOf(util.ToInt64(v) - 1)
			}
			if env.Set(aLHS.Lit, v) != nil {
				env.Define(aLHS.Lit, v)
			}
			return v, nil
		}
	}

	binop := &BinOp{Left: expr.Left, Operator: expr.Operator[0:1], Right: expr.Right}
	v, err := binop.Invoke(env)
	if err != nil {
		return v, err
	}

	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return expr.Left.Assign(v, env)
}

// Assign a value to the expression and return it.
func (expr *Assoc) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
