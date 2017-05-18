package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Break defines a break statement.
type Break struct {
	interpreter.PosImpl
}

func (stmt *Break) String() string {
	return "break"
}

// Execute the statement.
func (stmt *Break) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.ErrBreak
}
