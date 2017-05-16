package goblin

import "reflect"

// ParenExpr defines a parent block expression.
type ParenExpr struct {
	PosImpl
	SubExpr Expr
}

// Invoke the expression and return a result.
func (expr *ParenExpr) Invoke(env *Env) (reflect.Value, error) {
	v, err := expr.SubExpr.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	return v, nil
}

// Assign a value to the expression and return it.
func (expr *ParenExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
