package expression

import (
	"reflect"
	"strings"

	"github.com/richardwilkes/goblin"
)

// Ident defines identifier expression.
type Ident struct {
	goblin.PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *Ident) Invoke(env *goblin.Env) (reflect.Value, error) {
	return env.Get(expr.Lit)
}

// Assign a value to the expression and return it.
func (expr *Ident) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	if env.Set(expr.Lit, rv) != nil {
		if strings.Contains(expr.Lit, ".") {
			return goblin.NilValue, goblin.NewErrorf(expr, "Undefined symbol '%s'", expr.Lit)
		}
		env.Define(expr.Lit, rv)
	}
	return rv, nil
}
