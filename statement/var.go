package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Variable defines a variable definition statement.
type Variable struct {
	ast.PosImpl
	Names []string
	Exprs []ast.Expr
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
func (stmt *Variable) Execute(scope ast.Scope) (reflect.Value, error) {
	rvs := make([]reflect.Value, 0, len(stmt.Exprs))
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(scope)
		if err != nil {
			return rv, ast.NewError(expr, err)
		}
		rvs = append(rvs, rv)
	}
	result := make([]interface{}, 0, len(rvs))
	for i, name := range stmt.Names {
		if i < len(rvs) {
			scope.Define(name, rvs[i])
			result = append(result, rvs[i].Interface())
		}
	}
	return reflect.ValueOf(result), nil
}
