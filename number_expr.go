package goblin

import (
	"reflect"
	"strconv"
	"strings"
)

// NumberExpr defines a number expression.
type NumberExpr struct {
	PosImpl
	Lit string
}

// Invoke the expression and return a result.
func (expr *NumberExpr) Invoke(env *Env) (reflect.Value, error) {
	if strings.Contains(expr.Lit, ".") || strings.Contains(expr.Lit, "e") {
		v, err := strconv.ParseFloat(expr.Lit, 64)
		if err != nil {
			return NilValue, NewError(expr, err)
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
		return NilValue, NewError(expr, err)
	}
	return reflect.ValueOf(i), nil
}

// Assign a value to the expression and return it.
func (expr *NumberExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
