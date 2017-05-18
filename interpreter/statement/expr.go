package statement

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Expression defines an expression statement.
type Expression struct {
	interpreter.PosImpl
	Expr interpreter.Expr
}

func (stmt *Expression) String() string {
	return fmt.Sprint(stmt.Expr)
}

// Execute the statement.
func (stmt *Expression) Execute(env *interpreter.Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(stmt, err)
	}
	return rv, nil
}
