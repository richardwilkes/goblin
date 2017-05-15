package goblin

import "reflect"

// AssocExpr defines an operator association expression.
type AssocExpr struct {
	PosImpl
	LHS      Expr
	Operator string
	RHS      Expr
}

// Invoke the expression and return a result.
func (expr *AssocExpr) Invoke(env *Env) (reflect.Value, error) {
	switch expr.Operator {
	case "++":
		if aLHS, ok := expr.LHS.(*IdentExpr); ok {
			v, err := env.Get(aLHS.Lit)
			if err != nil {
				return v, err
			}
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(toFloat64(v) + 1.0)
			} else {
				v = reflect.ValueOf(toInt64(v) + 1)
			}
			if env.Set(aLHS.Lit, v) != nil {
				env.Define(aLHS.Lit, v)
			}
			return v, nil
		}
	case "--":
		if aLHS, ok := expr.LHS.(*IdentExpr); ok {
			v, err := env.Get(aLHS.Lit)
			if err != nil {
				return v, err
			}
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(toFloat64(v) - 1.0)
			} else {
				v = reflect.ValueOf(toInt64(v) - 1)
			}
			if env.Set(aLHS.Lit, v) != nil {
				env.Define(aLHS.Lit, v)
			}
			return v, nil
		}
	}

	binop := &BinOpExpr{LHS: expr.LHS, Operator: expr.Operator[0:1], RHS: expr.RHS}
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
func (expr *AssocExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
