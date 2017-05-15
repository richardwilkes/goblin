package goblin

import (
	"reflect"
	"strings"
)

// IdentExpr defines identifier expression.
type IdentExpr struct {
	PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *IdentExpr) Invoke(env *Env) (reflect.Value, error) {
	return env.Get(expr.Lit)
}

// Assign a value to the expression and return it.
func (expr *IdentExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	if env.Set(expr.Lit, rv) != nil {
		if strings.Contains(expr.Lit, ".") {
			return NilValue, NewErrorf(expr, "Undefined symbol '%s'", expr.Lit)
		}
		env.Define(expr.Lit, rv)
	}
	return rv, nil
}
