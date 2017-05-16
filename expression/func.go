package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Func defines a function expression.
type Func struct {
	goblin.PosImpl
	Name   string
	Stmts  []goblin.Stmt
	Args   []string
	VarArg bool
}

// Invoke the expression and return a result.
func (expr *Func) Invoke(env *goblin.Env) (reflect.Value, error) {
	f := reflect.ValueOf(func(fe *Func, env *goblin.Env) goblin.Func {
		return func(args ...reflect.Value) (reflect.Value, error) {
			if !fe.VarArg {
				if len(args) != len(fe.Args) {
					return goblin.NilValue, goblin.NewStringError(fe, fmt.Sprintf("Expecting %d arguments, got %d", len(fe.Args), len(args)))
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
			if err == goblin.ErrReturn {
				err = nil
				var ok bool
				if rr, ok = rr.Interface().(reflect.Value); !ok {
					rr = goblin.NilValue
				}
			}
			return rr, err
		}
	}(expr, env))
	env.Define(expr.Name, f)
	return f, nil
}

// Assign a value to the expression and return it.
func (expr *Func) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}