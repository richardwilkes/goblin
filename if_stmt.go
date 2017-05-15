package goblin

import (
	"reflect"

	"github.com/richardwilkes/goblin/util"
)

// IfStmt defines an if/else statement.
type IfStmt struct {
	PosImpl
	If     Expr
	Then   []Stmt
	ElseIf []Stmt
	Else   []Stmt
}

// Execute the statement.
func (stmt *IfStmt) Execute(env *Env) (reflect.Value, error) {
	// If
	rv, err := stmt.If.Invoke(env)
	if err != nil {
		return rv, NewError(stmt, err)
	}
	if util.ToBool(rv) {
		// Then
		newEnv := env.NewEnv()
		defer newEnv.Destroy()
		rv, err = newEnv.Run(stmt.Then)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		return rv, nil
	}
	done := false
	if len(stmt.ElseIf) > 0 {
		for _, stmt := range stmt.ElseIf {
			stmtIf, ok := stmt.(*IfStmt)
			if !ok {
				return NilValue, NewError(stmt, ErrBadSyntax)
			}
			// Else If
			rv, err = stmtIf.If.Invoke(env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if !util.ToBool(rv) {
				continue
			}
			// Else If Then
			done = true
			rv, err = env.Run(stmtIf.Then)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			break
		}
	}
	if !done && len(stmt.Else) > 0 {
		// Else
		newEnv := env.NewEnv()
		defer newEnv.Destroy()
		rv, err = newEnv.Run(stmt.Else)
		if err != nil {
			return rv, NewError(stmt, err)
		}
	}
	return rv, nil
}
