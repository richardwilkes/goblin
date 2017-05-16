package expression

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// MakeArray defines a make array expression.
type MakeArray struct {
	goblin.PosImpl
	Type    string
	LenExpr goblin.Expr
	CapExpr goblin.Expr
}

// Invoke the expression and return a result.
func (expr *MakeArray) Invoke(env *goblin.Env) (reflect.Value, error) {
	typ, err := env.Type(expr.Type)
	if err != nil {
		return goblin.NilValue, err
	}
	var alen int
	if expr.LenExpr != nil {
		rv, lerr := expr.LenExpr.Invoke(env)
		if lerr != nil {
			return goblin.NilValue, lerr
		}
		alen = int(util.ToInt64(rv))
	}
	var acap int
	if expr.CapExpr != nil {
		rv, lerr := expr.CapExpr.Invoke(env)
		if lerr != nil {
			return goblin.NilValue, lerr
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
func (expr *MakeArray) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
