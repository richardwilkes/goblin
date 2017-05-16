package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// BreakStmt defines a break statement.
type BreakStmt struct {
	goblin.PosImpl
}

// Execute the statement.
func (stmt *BreakStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.ErrBreak
}
