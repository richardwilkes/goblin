package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// If defines an if/else statement.
type If struct {
	goblin.PosImpl
	If     goblin.Expr
	Then   []goblin.Stmt
	ElseIf []goblin.Stmt
	Else   []goblin.Stmt
}

// Execute the statement.
func (stmt *If) Execute(env *goblin.Env) (reflect.Value, error) {
	// If
	rv, err := stmt.If.Invoke(env)
	if err != nil {
		return rv, goblin.NewError(stmt, err)
	}
	if util.ToBool(rv) {
		// Then
		newEnv := env.NewEnv()
		defer newEnv.Destroy()
		rv, err = newEnv.Run(stmt.Then)
		if err != nil {
			return rv, goblin.NewError(stmt, err)
		}
		return rv, nil
	}
	done := false
	if len(stmt.ElseIf) > 0 {
		for _, stmt := range stmt.ElseIf {
			stmtIf, ok := stmt.(*If)
			if !ok {
				return goblin.NilValue, goblin.NewError(stmt, goblin.ErrBadSyntax)
			}
			// Else If
			rv, err = stmtIf.If.Invoke(env)
			if err != nil {
				return rv, goblin.NewError(stmt, err)
			}
			if !util.ToBool(rv) {
				continue
			}
			// Else If Then
			done = true
			rv, err = env.Run(stmtIf.Then)
			if err != nil {
				return rv, goblin.NewError(stmt, err)
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
			return rv, goblin.NewError(stmt, err)
		}
	}
	return rv, nil
}
