package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Continue defines the continue statement.
type Continue struct {
	interpreter.PosImpl
}

// Execute the statement.
func (stmt *Continue) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.ErrContinue
}
