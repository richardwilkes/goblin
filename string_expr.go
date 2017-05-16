package goblin

import "reflect"

// StringExpr defines a string expression.
type StringExpr struct {
	PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *StringExpr) Invoke(env *Env) (reflect.Value, error) {
	return reflect.ValueOf(expr.Lit), nil
}

// Assign a value to the expression and return it.
func (expr *StringExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
