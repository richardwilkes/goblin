package statement

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Throw defines the throw statement.
type Throw struct {
	goblin.PosImpl
	Expr goblin.Expr
}

// Execute the statement.
func (stmt *Throw) Execute(env *goblin.Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(stmt, err)
	}
	if !rv.IsValid() {
		return goblin.NilValue, goblin.NewError(stmt, err)
	}
	return rv, goblin.NewStringError(stmt, fmt.Sprint(rv.Interface()))
}
