package goblin

import "reflect"

// BreakStmt defines a break statement.
type BreakStmt struct {
	PosImpl
}

// Execute the statement.
func (stmt *BreakStmt) Execute(env *Env) (reflect.Value, error) {
	return NilValue, ErrBreak
}
