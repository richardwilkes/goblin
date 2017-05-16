package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Expression defines an expression statement.
type Expression struct {
	interpreter.PosImpl
	Expr interpreter.Expr
}

// Execute the statement.
func (stmt *Expression) Execute(env *interpreter.Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(stmt, err)
	}
	return rv, nil
}
