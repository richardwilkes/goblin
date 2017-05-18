package expression

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/richardwilkes/goblin/interpreter"
)

// Number defines a number expression.
type Number struct {
	interpreter.PosImpl
	Lit      string
	value    reflect.Value
	err      error
	resolved bool
}

func (expr *Number) String() string {
	return expr.Lit
}

// Invoke the expression and return a result.
func (expr *Number) Invoke(env *interpreter.Env) (reflect.Value, error) {
	if !expr.resolved {
		expr.resolved = true
		var err error
		if strings.Contains(expr.Lit, ".") || strings.Contains(expr.Lit, "e") {
			var f float64
			f, err = strconv.ParseFloat(expr.Lit, 64)
			if err == nil {
				expr.value = reflect.ValueOf(f)
			}
		} else {
			var i int64
			if strings.HasPrefix(expr.Lit, "0x") {
				i, err = strconv.ParseInt(expr.Lit[2:], 16, 64)
			} else {
				i, err = strconv.ParseInt(expr.Lit, 10, 64)
			}
			if err == nil {
				expr.value = reflect.ValueOf(i)
			}
		}
		if err != nil {
			expr.value = interpreter.NilValue
			expr.err = interpreter.NewError(expr, err)
		}
	}
	return expr.value, expr.err
}

// Assign a value to the expression and return it.
func (expr *Number) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
