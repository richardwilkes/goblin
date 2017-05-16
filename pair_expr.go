package goblin

import "reflect"

// PairExpr defines a map key/value pair expression.
type PairExpr struct {
	PosImpl
	Key   string
	Value Expr
}

// Invoke the expression and return a result.
func (expr *PairExpr) Invoke(env *Env) (reflect.Value, error) {
	return NilValue, NewStringError(expr, "Not invokable")
}

// Assign a value to the expression and return it.
func (expr *PairExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
