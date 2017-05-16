package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Continue defines the continue statement.
type Continue struct {
	goblin.PosImpl
}

// Execute the statement.
func (stmt *Continue) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.ErrContinue
}
