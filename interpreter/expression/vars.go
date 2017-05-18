package expression

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Vars defines an expression that defines multiple variables.
type Vars struct {
	interpreter.PosImpl
	Left     []interpreter.Expr
	Operator string
	Right    []interpreter.Expr
}

func (expr *Vars) String() string {
	var buffer bytes.Buffer
	for i, one := range expr.Left {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	buffer.WriteString(" ")
	buffer.WriteString(expr.Operator)
	buffer.WriteString(" ")
	for i, one := range expr.Right {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Vars) Invoke(env *interpreter.Env) (reflect.Value, error) {
	rv := interpreter.NilValue
	var err error
	vs := []interface{}{}
	for _, Right := range expr.Right {
		rv, err = Right.Invoke(env)
		if err != nil {
			return rv, interpreter.NewError(Right, err)
		}
		if rv == interpreter.NilValue {
			vs = append(vs, nil)
		} else if rv.IsValid() && rv.CanInterface() {
			vs = append(vs, rv.Interface())
		} else {
			vs = append(vs, nil)
		}
	}
	rvs := reflect.ValueOf(vs)
	for i, Left := range expr.Left {
		if i >= rvs.Len() {
			break
		}
		v := rvs.Index(i)
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		_, err = Left.Assign(v, env)
		if err != nil {
			return rvs, interpreter.NewError(Left, err)
		}
	}
	return rvs, nil
}

// Assign a value to the expression and return it.
func (expr *Vars) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
