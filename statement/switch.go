package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// Switch defines a switch statement.
type Switch struct {
	goblin.PosImpl
	Expr  goblin.Expr
	Cases []goblin.Stmt
}

// Execute the statement.
func (stmt *Switch) Execute(env *goblin.Env) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(stmt, err)
	}
	done := false
	var defaultStmt *Default
	for _, ss := range stmt.Cases {
		if ssd, ok := ss.(*Default); ok {
			defaultStmt = ssd
			continue
		}
		caseStmt, ok := ss.(*Case)
		if !ok {
			return goblin.NilValue, goblin.NewError(stmt, goblin.ErrBadSyntax)
		}
		cv, lerr := caseStmt.Expr.Invoke(env)
		if lerr != nil {
			return rv, goblin.NewError(stmt, lerr)
		}
		if !util.Equal(rv, cv) {
			continue
		}
		rv, err = env.Run(caseStmt.Stmts)
		if err != nil {
			return rv, goblin.NewError(stmt, err)
		}
		done = true
		break
	}
	if !done && defaultStmt != nil {
		rv, err = env.Run(defaultStmt.Stmts)
		if err != nil {
			return rv, goblin.NewError(stmt, err)
		}
	}
	return rv, nil
}
