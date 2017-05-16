package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Try defines the try/catch/finally statement.
type Try struct {
	goblin.PosImpl
	Try     []goblin.Stmt
	Var     string
	Catch   []goblin.Stmt
	Finally []goblin.Stmt
}

// Execute the statement.
func (stmt *Try) Execute(env *goblin.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	_, err := newEnv.Run(stmt.Try)
	if err != nil {
		// Catch
		cenv := env.NewEnv()
		defer cenv.Destroy()
		if stmt.Var != "" {
			cenv.Define(stmt.Var, reflect.ValueOf(err))
		}
		_, e1 := cenv.Run(stmt.Catch)
		if e1 != nil {
			err = goblin.NewError(stmt.Catch[0], e1)
		} else {
			err = nil
		}
	}
	if len(stmt.Finally) > 0 {
		// Finally
		fenv := env.NewEnv()
		defer fenv.Destroy()
		_, e2 := fenv.Run(stmt.Finally)
		if e2 != nil {
			err = goblin.NewError(stmt.Finally[0], e2)
		}
	}
	return goblin.NilValue, goblin.NewError(stmt, err)
}
