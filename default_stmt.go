package goblin

import "reflect"

// DefaultStmt defines the default case statement.
type DefaultStmt struct {
	PosImpl
	Stmts []Stmt
}

// Execute the statement.
func (stmt *DefaultStmt) Execute(env *Env) (reflect.Value, error) {
	return NilValue, newCannotExecuteError(stmt)
}
