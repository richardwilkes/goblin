package goblin

import "reflect"

// TernaryOpExpr defines a ternary operator expression.
type TernaryOpExpr struct {
	PosImpl
	Expr Expr
	LHS  Expr
	RHS  Expr
}

// Invoke the expression and return a result.
func (expr *TernaryOpExpr) Invoke(env *Env) (reflect.Value, error) {
	rv, err := expr.Expr.Invoke(env)
	if err != nil {
		return rv, NewError(expr, err)
	}
	var choice Expr
	if toBool(rv) {
		choice = expr.LHS
	} else {
		choice = expr.RHS
	}
	rv, err = choice.Invoke(env)
	if err != nil {
		return rv, NewError(choice, err)
	}
	return rv, nil
}

// Assign a value to the expression and return it.
func (expr *TernaryOpExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
