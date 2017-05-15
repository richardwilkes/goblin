package goblin

import "reflect"

// LetExpr defines an expression to define a variable.
type LetExpr struct {
	PosImpl
	LHS Expr
	RHS Expr
}

// Invoke the expression and return a result.
func (expr *LetExpr) Invoke(env *Env) (reflect.Value, error) {
	rv, err := expr.RHS.Invoke(env)
	if err != nil {
		return rv, NewError(expr, err)
	}
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return expr.LHS.Assign(rv, env)
}

// Assign a value to the expression and return it.
func (expr *LetExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
