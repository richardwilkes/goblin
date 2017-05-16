package goblin

import "reflect"

// MapExpr defines a map expression.
type MapExpr struct {
	PosImpl
	MapExpr map[string]Expr
}

// Invoke the expression and return a result.
func (expr *MapExpr) Invoke(env *Env) (reflect.Value, error) {
	m := make(map[string]interface{})
	for k, e := range expr.MapExpr {
		v, err := e.Invoke(env)
		if err != nil {
			return v, NewError(e, err)
		}
		m[k] = v.Interface()
	}
	return reflect.ValueOf(m), nil
}

// Assign a value to the expression and return it.
func (expr *MapExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
