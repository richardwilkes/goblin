package statement

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Throw defines the throw statement.
type Throw struct {
	interpreter.PosImpl
	Expr interpreter.Expr
}

func (stmt *Throw) String() string {
	return fmt.Sprintf("throw %v", stmt.Expr)
}

// Execute the statement.
func (stmt *Throw) Execute(env *interpreter.Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(stmt, err)
	}
	if !rv.IsValid() {
		return interpreter.NilValue, interpreter.NewError(stmt, err)
	}
	return rv, interpreter.NewStringError(stmt, fmt.Sprint(rv.Interface()))
}
