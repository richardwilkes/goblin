package expression

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// MakeArray defines a make array expression.
type MakeArray struct {
	interpreter.PosImpl
	Type    string
	LenExpr interpreter.Expr
	CapExpr interpreter.Expr
}

func (expr *MakeArray) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("make([]")
	buffer.WriteString(expr.Type)
	if expr.LenExpr != nil {
		fmt.Fprintf(&buffer, ",%v", expr.LenExpr)
	}
	if expr.CapExpr != nil {
		fmt.Fprintf(&buffer, ",%v", expr.CapExpr)
	}
	buffer.WriteString(")")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *MakeArray) Invoke(env *interpreter.Env) (reflect.Value, error) {
	typ, err := env.Type(expr.Type)
	if err != nil {
		return interpreter.NilValue, err
	}
	var alen int
	if expr.LenExpr != nil {
		rv, lerr := expr.LenExpr.Invoke(env)
		if lerr != nil {
			return interpreter.NilValue, lerr
		}
		alen = int(util.ToInt64(rv))
	}
	var acap int
	if expr.CapExpr != nil {
		rv, lerr := expr.CapExpr.Invoke(env)
		if lerr != nil {
			return interpreter.NilValue, lerr
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
func (expr *MakeArray) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
