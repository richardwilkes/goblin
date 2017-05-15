package goblin

import "reflect"

// ForStmt defines a for statement.
type ForStmt struct {
	PosImpl
	Var   string
	Value Expr
	Stmts []Stmt
}

// Execute the statement.
func (stmt *ForStmt) Execute(env *Env) (reflect.Value, error) {
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
				if err == ErrBreak {
					break
				}
				if err == ErrContinue {
					continue
				}
				if err == ErrReturn {
					return rv, err
				}
				return rv, NewError(stmt, err)
			}
		}
		return NilValue, nil
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
				if err == ErrBreak {
					break
				}
				if err == ErrContinue {
					continue
				}
				if err == ErrReturn {
					return rv, err
				}
				return rv, NewError(stmt, err)
			}
		}
		return NilValue, nil
	} else {
		return NilValue, NewStringError(stmt, "Invalid operation for non-array value")
	}
}
