package goblin

import "reflect"

// ContinueStmt defines the continue statement.
type ContinueStmt struct {
	PosImpl
}

// Execute the statement.
func (stmt *ContinueStmt) Execute(env *Env) (reflect.Value, error) {
	return NilValue, ErrContinue
}
