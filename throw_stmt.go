package goblin

import (
	"fmt"
	"reflect"
)

// ThrowStmt defines the throw statement.
type ThrowStmt struct {
	PosImpl
	Expr Expr
}

// Execute the statement.
func (stmt *ThrowStmt) Execute(env *Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, NewError(stmt, err)
	}
	if !rv.IsValid() {
		return NilValue, NewError(stmt, err)
	}
	return rv, NewStringError(stmt, fmt.Sprint(rv.Interface()))
}
