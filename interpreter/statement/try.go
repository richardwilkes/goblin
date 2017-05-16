package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Try defines the try/catch/finally statement.
type Try struct {
	interpreter.PosImpl
	Try     []interpreter.Stmt
	Var     string
	Catch   []interpreter.Stmt
	Finally []interpreter.Stmt
}

// Execute the statement.
func (stmt *Try) Execute(env *interpreter.Env) (reflect.Value, error) {
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
			err = interpreter.NewError(stmt.Catch[0], e1)
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
			err = interpreter.NewError(stmt.Finally[0], e2)
		}
	}
	return interpreter.NilValue, interpreter.NewError(stmt, err)
}
