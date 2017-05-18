package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Variable defines a variable definition statement.
type Variable struct {
	interpreter.PosImpl
	Names []string
	Exprs []interpreter.Expr
}

func (stmt *Variable) String() string {
	var buffer bytes.Buffer
	for i, name := range stmt.Names {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(name)
	}
	buffer.WriteString(" = ")
	for i, one := range stmt.Exprs {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Variable) Execute(env *interpreter.Env) (reflect.Value, error) {
	rvs := make([]reflect.Value, 0, len(stmt.Exprs))
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(env)
		if err != nil {
			return rv, interpreter.NewError(expr, err)
		}
		rvs = append(rvs, rv)
	}
	result := make([]interface{}, 0, len(rvs))
	for i, name := range stmt.Names {
		if i < len(rvs) {
			env.Define(name, rvs[i])
			result = append(result, rvs[i].Interface())
		}
	}
	return reflect.ValueOf(result), nil
}
