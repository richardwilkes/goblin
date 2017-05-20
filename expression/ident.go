package expression

import (
	"reflect"
	"strings"

	"github.com/richardwilkes/goblin/ast"
)

// Ident defines identifier expression.
type Ident struct {
	ast.PosImpl
	Lit string
}

func (expr *Ident) String() string {
	return expr.Lit
}

// Invoke the expression and return a result.
func (expr *Ident) Invoke(scope ast.Scope) (reflect.Value, error) {
	return scope.Get(expr.Lit)
}

// Assign a value to the expression and return it.
func (expr *Ident) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	if scope.Set(expr.Lit, rv) != nil {
		if strings.Contains(expr.Lit, ".") {
			return ast.NilValue, ast.NewErrorf(expr, "Undefined symbol '%s'", expr.Lit)
		}
		scope.Define(expr.Lit, rv)
	}
	return rv, nil
}
