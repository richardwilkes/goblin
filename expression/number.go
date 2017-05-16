package expression

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/richardwilkes/goblin"
)

// Number defines a number expression.
type Number struct {
	goblin.PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *Number) Invoke(env *goblin.Env) (reflect.Value, error) {
	if strings.Contains(expr.Lit, ".") || strings.Contains(expr.Lit, "e") {
		v, err := strconv.ParseFloat(expr.Lit, 64)
		if err != nil {
			return goblin.NilValue, goblin.NewError(expr, err)
		}
		return reflect.ValueOf(v), nil
	}
	var i int64
	var err error
	if strings.HasPrefix(expr.Lit, "0x") {
		i, err = strconv.ParseInt(expr.Lit[2:], 16, 64)
	} else {
		i, err = strconv.ParseInt(expr.Lit, 10, 64)
	}
	if err != nil {
		return goblin.NilValue, goblin.NewError(expr, err)
	}
	return reflect.ValueOf(i), nil
}

// Assign a value to the expression and return it.
func (expr *Number) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
