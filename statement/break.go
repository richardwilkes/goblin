package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Break defines a break statement.
type Break struct {
	goblin.PosImpl
}

// Execute the statement.
func (stmt *Break) Execute(env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.ErrBreak
}
