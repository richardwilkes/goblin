package goblin

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/util"
)

// MakeArrayExpr defines a make array expression.
type MakeArrayExpr struct {
	PosImpl
	Type    string
	LenExpr Expr
	CapExpr Expr
}

// Invoke the expression and return a result.
func (expr *MakeArrayExpr) Invoke(env *Env) (reflect.Value, error) {
	typ, err := env.Type(expr.Type)
	if err != nil {
		return NilValue, err
	}
	var alen int
	if expr.LenExpr != nil {
		rv, lerr := expr.LenExpr.Invoke(env)
		if lerr != nil {
			return NilValue, lerr
		}
		alen = int(util.ToInt64(rv))
	}
	var acap int
	if expr.CapExpr != nil {
		rv, lerr := expr.CapExpr.Invoke(env)
		if lerr != nil {
			return NilValue, lerr
		}
		acap = int(util.ToInt64(rv))
	} else {
		acap = alen
	}
	return func() (reflect.Value, error) {
		defer func() {
			if ex := recover(); ex != nil {
				if e, ok := ex.(error); ok {
					err = e
				} else {
					err = errors.New(fmt.Sprint(ex))
				}
			}
		}()
		return reflect.MakeSlice(reflect.SliceOf(typ), alen, acap), nil
	}()
}

// Assign a value to the expression and return it.
func (expr *MakeArrayExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
