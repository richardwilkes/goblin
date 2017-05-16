package goblin

import "reflect"

// AnonCallExpr defines an anonymous calling expression, e.g. func(){}().
type AnonCallExpr struct {
	PosImpl
	Expr     Expr
	SubExprs []Expr
	VarArg   bool
}

// Invoke the expression and return a result.
func (expr *AnonCallExpr) Invoke(env *Env) (reflect.Value, error) {
	f, err := expr.Expr.Invoke(env)
	if err != nil {
		return f, NewError(expr, err)
	}
	if f.Kind() == reflect.Interface {
		f = f.Elem()
	}
	if f.Kind() != reflect.Func {
		return f, NewStringError(expr, "Unknown function")
	}
	call := &CallExpr{Func: f, SubExprs: expr.SubExprs, VarArg: expr.VarArg}
	return call.Invoke(env)
}

// Assign a value to the expression and return it.
func (expr *AnonCallExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
