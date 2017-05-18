package expression

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Map defines a map expression.
type Map struct {
	interpreter.PosImpl
	Map map[string]interpreter.Expr
}

func (expr *Map) String() string {
	keys := make([]string, 0, len(expr.Map))
	for k := range expr.Map {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for i, k := range keys {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(util.QuotedString(k))
		buffer.WriteString(": ")
		fmt.Fprint(&buffer, expr.Map[k])
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Map) Invoke(env *interpreter.Env) (reflect.Value, error) {
	m := make(map[string]interface{})
	for k, e := range expr.Map {
		v, err := e.Invoke(env)
		if err != nil {
			return v, interpreter.NewError(e, err)
		}
		m[k] = v.Interface()
	}
	return reflect.ValueOf(m), nil
}

// Assign a value to the expression and return it.
func (expr *Map) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
