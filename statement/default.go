package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// DefaultStmt defines the default case statement.
type DefaultStmt struct {
	goblin.PosImpl
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *DefaultStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewCannotExecuteError(stmt)
}
