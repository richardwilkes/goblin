package expression

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Func defines a function expression.
type Func struct {
	ast.PosImpl
	Name   string
	Stmts  []ast.Stmt
	Args   []string
	VarArg bool
}

func (expr *Func) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("func")
	if expr.Name != "" {
		buffer.WriteString(" ")
		buffer.WriteString(expr.Name)
	}
	buffer.WriteString("(")
	for i, arg := range expr.Args {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(arg)
	}
	buffer.WriteString(") {")
	if len(expr.Stmts) > 0 {
		prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
		for _, stmt := range expr.Stmts {
			fmt.Fprintf(prefixer, "\n%v", stmt)
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Func) Invoke(scope ast.Scope) (reflect.Value, error) {
	f := reflect.ValueOf(func(fe *Func, scope ast.Scope) ast.Func {
		return func(args ...reflect.Value) (reflect.Value, error) {
			if !fe.VarArg {
				if len(args) != len(fe.Args) {
					return ast.NilValue, ast.NewStringError(fe, fmt.Sprintf("Expecting %d arguments, got %d", len(fe.Args), len(args)))
				}
			}
			newScope := scope.NewScope()
			if fe.VarArg {
				newScope.Define(fe.Args[0], reflect.ValueOf(args))
			} else {
				for i, arg := range fe.Args {
					newScope.Define(arg, args[i])
				}
			}
			rr, err := newScope.Run(fe.Stmts)
			if err == ast.ErrReturn {
				err = nil
			}
			return rr, err
		}
	}(expr, scope))
	scope.Define(expr.Name, f)
	return f, nil
}

// Assign a value to the expression and return it.
func (expr *Func) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}