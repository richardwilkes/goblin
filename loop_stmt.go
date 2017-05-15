package goblin

import "reflect"

// LoopStmt defines a loop statement.
type LoopStmt struct {
	PosImpl
	Expr  Expr
	Stmts []Stmt
}

// Execute the statement.
func (stmt *LoopStmt) Execute(env *Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	for {
		if stmt.Expr != nil {
			ev, ee := stmt.Expr.Invoke(newEnv)
			if ee != nil {
				return ev, ee
			}
			if !toBool(ev) {
				break
			}
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
	}
	return NilValue, nil
}
