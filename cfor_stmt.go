package goblin

import (
	"reflect"

	"github.com/richardwilkes/goblin/util"
)

// CForStmt defines a C-style "for (;;)" statement.
type CForStmt struct {
	PosImpl
	Expr1 Expr
	Expr2 Expr
	Expr3 Expr
	Stmts []Stmt
}

// Execute the statement.
func (stmt *CForStmt) Execute(env *Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	_, err := stmt.Expr1.Invoke(newEnv)
	if err != nil {
		return NilValue, err
	}
	for {
		fb, err := stmt.Expr2.Invoke(newEnv)
		if err != nil {
			return NilValue, err
		}
		if !util.ToBool(fb) {
			break
		}

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
		_, err = stmt.Expr3.Invoke(newEnv)
		if err != nil {
			return NilValue, err
		}
	}
	return NilValue, nil
}
