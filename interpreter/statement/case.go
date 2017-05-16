package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Case defines a case statement.
type Case struct {
	interpreter.PosImpl
	Expr  interpreter.Expr
	Stmts []interpreter.Stmt
}

// Execute the statement.
func (stmt *Case) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewCannotExecuteError(stmt)
}
