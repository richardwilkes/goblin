package goblin

import "reflect"

// ExprStmt defines an expression statement.
type ExprStmt struct {
	PosImpl
	Expr Expr
}

// Execute the statement.
func (stmt *ExprStmt) Execute(env *Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, NewError(stmt, err)
	}
	return rv, nil
}
