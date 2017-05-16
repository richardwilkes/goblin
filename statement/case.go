package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// CaseStmt defines a case statement.
type CaseStmt struct {
	goblin.PosImpl
	Expr  goblin.Expr
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *CaseStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewCannotExecuteError(stmt)
}
