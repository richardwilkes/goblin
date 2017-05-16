package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Default defines the default case statement.
type Default struct {
	interpreter.PosImpl
	Stmts []interpreter.Stmt
}

// Execute the statement.
func (stmt *Default) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewCannotExecuteError(stmt)
}
