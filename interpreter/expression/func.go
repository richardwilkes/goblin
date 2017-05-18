package expression

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/interpreter"
)

// Func defines a function expression.
type Func struct {
	interpreter.PosImpl
	Name   string
	Stmts  []interpreter.Stmt
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
		for _, stmt := range expr.Stmts {
			buffer.WriteString("\n    ")
			fmt.Fprint(&buffer, stmt)
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Func) Invoke(env *interpreter.Env) (reflect.Value, error) {
	f := reflect.ValueOf(func(fe *Func, env *interpreter.Env) interpreter.Func {
		return func(args ...reflect.Value) (reflect.Value, error) {
			if !fe.VarArg {
				if len(args) != len(fe.Args) {
					return interpreter.NilValue, interpreter.NewStringError(fe, fmt.Sprintf("Expecting %d arguments, got %d", len(fe.Args), len(args)))
				}
			}
			newEnv := env.NewEnv()
			if fe.VarArg {
				newEnv.Define(fe.Args[0], reflect.ValueOf(args))
			} else {
				for i, arg := range fe.Args {
					newEnv.Define(arg, args[i])
				}
			}
			rr, err := newEnv.Run(fe.Stmts)
			if err == interpreter.ErrReturn {
				err = nil
			}
			return rr, err
		}
	}(expr, env))
	env.Define(expr.Name, f)
	return f, nil
}

// Assign a value to the expression and return it.
func (expr *Func) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
