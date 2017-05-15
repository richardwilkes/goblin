package goblin

import "reflect"

// ConstExpr defines a constant.
type ConstExpr struct {
	PosImpl
	Value string
}

// Invoke the expression and return a result.
func (expr *ConstExpr) Invoke(env *Env) (reflect.Value, error) {
	switch expr.Value {
	case "true":
		return reflect.ValueOf(true), nil
	case "false":
		return reflect.ValueOf(false), nil
	}
	return reflect.ValueOf(nil), nil
}

// Assign a value to the expression and return it.
func (expr *ConstExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
