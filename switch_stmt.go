package goblin

import (
	"reflect"

	"github.com/richardwilkes/goblin/util"
)

// SwitchStmt defines a switch statement.
type SwitchStmt struct {
	PosImpl
	Expr  Expr
	Cases []Stmt
}

// Execute the statement.
func (stmt *SwitchStmt) Execute(env *Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, NewError(stmt, err)
	}
	done := false
	var defaultStmt *DefaultStmt
	for _, ss := range stmt.Cases {
		if ssd, ok := ss.(*DefaultStmt); ok {
			defaultStmt = ssd
			continue
		}
		caseStmt, ok := ss.(*CaseStmt)
		if !ok {
			return NilValue, NewError(stmt, ErrBadSyntax)
		}
		cv, lerr := caseStmt.Expr.Invoke(env)
		if lerr != nil {
			return rv, NewError(stmt, lerr)
		}
		if !util.Equal(rv, cv) {
			continue
		}
		rv, err = env.Run(caseStmt.Stmts)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		done = true
		break
	}
	if !done && defaultStmt != nil {
		rv, err = env.Run(defaultStmt.Stmts)
		if err != nil {
			return rv, NewError(stmt, err)
		}
	}
	return rv, nil
}
