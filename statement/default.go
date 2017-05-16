package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Default defines the default case statement.
type Default struct {
	goblin.PosImpl
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *Default) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewCannotExecuteError(stmt)
}
