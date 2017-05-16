package statement

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin"
)

// ThrowStmt defines the throw statement.
type ThrowStmt struct {
	goblin.PosImpl
	Expr goblin.Expr
}

// Execute the statement.
func (stmt *ThrowStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(stmt, err)
	}
	if !rv.IsValid() {
		return goblin.NilValue, goblin.NewError(stmt, err)
	}
	return rv, goblin.NewStringError(stmt, fmt.Sprint(rv.Interface()))
}
