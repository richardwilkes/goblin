package expression

import (
	"reflect"
	"strings"

	"github.com/richardwilkes/goblin/interpreter"
)

// Ident defines identifier expression.
type Ident struct {
	interpreter.PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *Ident) Invoke(env *interpreter.Env) (reflect.Value, error) {
	return env.Get(expr.Lit)
}

// Assign a value to the expression and return it.
func (expr *Ident) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	if env.Set(expr.Lit, rv) != nil {
		if strings.Contains(expr.Lit, ".") {
			return interpreter.NilValue, interpreter.NewErrorf(expr, "Undefined symbol '%s'", expr.Lit)
		}
		env.Define(expr.Lit, rv)
	}
	return rv, nil
}
