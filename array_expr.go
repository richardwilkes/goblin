package goblin

import "reflect"

// ArrayExpr defines an array expression.
type ArrayExpr struct {
	PosImpl
	Exprs []Expr
}

// Invoke the expression and return a result.
func (expr *ArrayExpr) Invoke(env *Env) (reflect.Value, error) {
	a := make([]interface{}, len(expr.Exprs))
	for i, e := range expr.Exprs {
		arg, err := e.Invoke(env)
		if err != nil {
			return arg, NewError(e, err)
		}
		a[i] = arg.Interface()
	}
	return reflect.ValueOf(a), nil
}

// Assign a value to the expression and return it.
func (expr *ArrayExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
