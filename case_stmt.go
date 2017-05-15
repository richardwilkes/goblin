package goblin

import "reflect"

// CaseStmt defines a case statement.
type CaseStmt struct {
	PosImpl
	Expr  Expr
	Stmts []Stmt
}

// Execute the statement.
func (stmt *CaseStmt) Execute(env *Env) (reflect.Value, error) {
	return NilValue, newCannotExecuteError(stmt)
}
