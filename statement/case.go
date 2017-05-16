package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Case defines a case statement.
type Case struct {
	goblin.PosImpl
	Expr  goblin.Expr
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *Case) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewCannotExecuteError(stmt)
}
