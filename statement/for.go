package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// For defines a for statement.
type For struct {
	goblin.PosImpl
	Var   string
	Value goblin.Expr
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *For) Execute(env *goblin.Env) (reflect.Value, error) {
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
				if err == goblin.ErrBreak {
					break
				}
				if err == goblin.ErrContinue {
					continue
				}
				if err == goblin.ErrReturn {
					return rv, err
				}
				return rv, goblin.NewError(stmt, err)
			}
		}
		return goblin.NilValue, nil
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
				if err == goblin.ErrBreak {
					break
				}
				if err == goblin.ErrContinue {
					continue
				}
				if err == goblin.ErrReturn {
					return rv, err
				}
				return rv, goblin.NewError(stmt, err)
			}
		}
		return goblin.NilValue, nil
	} else {
		return goblin.NilValue, goblin.NewStringError(stmt, "Invalid operation for non-array value")
	}
}
