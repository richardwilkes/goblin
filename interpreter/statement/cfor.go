package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// CFor defines a C-style "for (;;)" statement.
type CFor struct {
	interpreter.PosImpl
	Expr1 interpreter.Expr
	Expr2 interpreter.Expr
	Expr3 interpreter.Expr
	Stmts []interpreter.Stmt
}

// Execute the statement.
func (stmt *CFor) Execute(env *interpreter.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	_, err := stmt.Expr1.Invoke(newEnv)
	if err != nil {
		return interpreter.NilValue, err
	}
	for {
		fb, err := stmt.Expr2.Invoke(newEnv)
		if err != nil {
			return interpreter.NilValue, err
		}
		if !util.ToBool(fb) {
			break
		}

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
		_, err = stmt.Expr3.Invoke(newEnv)
		if err != nil {
			return interpreter.NilValue, err
		}
	}
	return interpreter.NilValue, nil
}