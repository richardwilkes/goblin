package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// Loop defines a loop statement.
type Loop struct {
	goblin.PosImpl
	Expr  goblin.Expr
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *Loop) Execute(env *goblin.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	for {
		if stmt.Expr != nil {
			ev, ee := stmt.Expr.Invoke(newEnv)
			if ee != nil {
				return ev, ee
			}
			if !util.ToBool(ev) {
				break
			}
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
	}
	return goblin.NilValue, nil
}
