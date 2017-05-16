package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Assoc defines an operator association expression.
type Assoc struct {
	interpreter.PosImpl
	LHS      interpreter.Expr
	Operator string
	RHS      interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Assoc) Invoke(env *interpreter.Env) (reflect.Value, error) {
	switch expr.Operator {
	case "++":
		if aLHS, ok := expr.LHS.(*Ident); ok {
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
		if aLHS, ok := expr.LHS.(*Ident); ok {
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

	binop := &BinOp{LHS: expr.LHS, Operator: expr.Operator[0:1], RHS: expr.RHS}
	v, err := binop.Invoke(env)
	if err != nil {
		return v, err
	}

	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return expr.LHS.Assign(v, env)
}

// Assign a value to the expression and return it.
func (expr *Assoc) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
