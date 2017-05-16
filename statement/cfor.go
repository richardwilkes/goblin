package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// CFor defines a C-style "for (;;)" statement.
type CFor struct {
	goblin.PosImpl
	Expr1 goblin.Expr
	Expr2 goblin.Expr
	Expr3 goblin.Expr
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *CFor) Execute(env *goblin.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	_, err := stmt.Expr1.Invoke(newEnv)
	if err != nil {
		return goblin.NilValue, err
	}
	for {
		fb, err := stmt.Expr2.Invoke(newEnv)
		if err != nil {
			return goblin.NilValue, err
		}
		if !util.ToBool(fb) {
			break
		}

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
		_, err = stmt.Expr3.Invoke(newEnv)
		if err != nil {
			return goblin.NilValue, err
		}
	}
	return goblin.NilValue, nil
}
