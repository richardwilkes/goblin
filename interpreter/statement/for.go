package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// For defines a for statement.
type For struct {
	interpreter.PosImpl
	Var   string
	Value interpreter.Expr
	Stmts []interpreter.Stmt
}

// Execute the statement.
func (stmt *For) Execute(env *interpreter.Env) (reflect.Value, error) {
	val, ee := stmt.Value.Invoke(env)
	if ee != nil {
		return val, ee
	}
	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		newEnv := env.NewEnv()
		defer newEnv.Destroy()

		for i := 0; i < val.Len(); i++ {
			iv := val.Index(i)
			if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
				iv = iv.Elem()
			}
			newEnv.Define(stmt.Var, iv)
			rv, err := newEnv.Run(stmt.Stmts)
			if err != nil {
				if err == interpreter.ErrBreak {
					break
				}
				if err == interpreter.ErrContinue {
					continue
				}
				if err == interpreter.ErrReturn {
					return rv, err
				}
				return rv, interpreter.NewError(stmt, err)
			}
		}
		return interpreter.NilValue, nil
	} else if val.Kind() == reflect.Chan {
		newEnv := env.NewEnv()
		defer newEnv.Destroy()

		for {
			iv, ok := val.Recv()
			if !ok {
				break
			}
			if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
				iv = iv.Elem()
			}
			newEnv.Define(stmt.Var, iv)
			rv, err := newEnv.Run(stmt.Stmts)
			if err != nil {
				if err == interpreter.ErrBreak {
					break
				}
				if err == interpreter.ErrContinue {
					continue
				}
				if err == interpreter.ErrReturn {
					return rv, err
				}
				return rv, interpreter.NewError(stmt, err)
			}
		}
		return interpreter.NilValue, nil
	} else {
		return interpreter.NilValue, interpreter.NewStringError(stmt, "Invalid operation for non-array value")
	}
}
