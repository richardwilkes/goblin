package expression

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/interpreter"
)

// Array defines an array expression.
type Array struct {
	interpreter.PosImpl
	Exprs []interpreter.Expr
}

func (expr *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, v := range expr.Exprs {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprint(&buffer, v)
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Array) Invoke(env *interpreter.Env) (reflect.Value, error) {
	a := make([]interface{}, len(expr.Exprs))
	for i, e := range expr.Exprs {
		arg, err := e.Invoke(env)
		if err != nil {
			return arg, interpreter.NewError(e, err)
		}
		a[i] = arg.Interface()
	}
	return reflect.ValueOf(a), nil
}

// Assign a value to the expression and return it.
func (expr *Array) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
