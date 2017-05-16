package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// ContinueStmt defines the continue statement.
type ContinueStmt struct {
	goblin.PosImpl
}

// Execute the statement.
func (stmt *ContinueStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.ErrContinue
}
