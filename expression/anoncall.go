package expression

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// AnonCall defines an anonymous calling expression, e.g. func(){}().
type AnonCall struct {
	ast.PosImpl
	Expr     ast.Expr
	SubExprs []ast.Expr
	VarArg   bool
}

func (expr *AnonCall) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%v(", expr.Expr)
	for i, arg := range expr.SubExprs {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprint(&buffer, arg)
	}
	if expr.VarArg {
		buffer.WriteString("...")
	}
	buffer.WriteString(")")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *AnonCall) Invoke(scope ast.Scope) (reflect.Value, error) {
	f, err := expr.Expr.Invoke(scope)
	if err != nil {
		return f, ast.NewError(expr, err)
	}
	if f.Kind() == reflect.Interface {
		f = f.Elem()
	}
	if f.Kind() != reflect.Func {
		return f, ast.NewStringError(expr, "Unknown function")
	}
	call := &Call{Func: f, SubExprs: expr.SubExprs, VarArg: expr.VarArg}
	return call.Invoke(scope)
}

// Assign a value to the expression and return it.
func (expr *AnonCall) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
