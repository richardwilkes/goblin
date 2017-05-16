package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// ExprStmt defines an expression statement.
type ExprStmt struct {
	goblin.PosImpl
	Expr goblin.Expr
}

// Execute the statement.
func (stmt *ExprStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(stmt, err)
	}
	return rv, nil
}
