package goblin

import (
	"fmt"
	"reflect"
)

// FuncExpr defines a function expression.
type FuncExpr struct {
	PosImpl
	Name   string
	Stmts  []Stmt
	Args   []string
	VarArg bool
}

// Invoke the expression and return a result.
func (expr *FuncExpr) Invoke(env *Env) (reflect.Value, error) {
	f := reflect.ValueOf(func(fe *FuncExpr, env *Env) Func {
		return func(args ...reflect.Value) (reflect.Value, error) {
			if !fe.VarArg {
				if len(args) != len(fe.Args) {
					return NilValue, NewStringError(fe, fmt.Sprintf("Expecting %d arguments, got %d", len(fe.Args), len(args)))
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
			if err == ErrReturn {
				err = nil
				var ok bool
				if rr, ok = rr.Interface().(reflect.Value); !ok {
					rr = NilValue
				}
			}
			return rr, err
		}
	}(expr, env))
	env.Define(expr.Name, f)
	return f, nil
}

// Assign a value to the expression and return it.
func (expr *FuncExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
