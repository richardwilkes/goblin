package goblin

import "reflect"

// MakeExpr defines a make expression.
type MakeExpr struct {
	PosImpl
	Type string
}

// Invoke the expression and return a result.
func (expr *MakeExpr) Invoke(env *Env) (reflect.Value, error) {
	rt, err := env.Type(expr.Type)
	if err != nil {
		return NilValue, NewError(expr, err)
	}
	if rt.Kind() == reflect.Map {
		return reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem())).Convert(rt), nil
	}
	return reflect.Zero(rt), nil
}

// Assign a value to the expression and return it.
func (expr *MakeExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
