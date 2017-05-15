package goblin

import "reflect"

// TryStmt defines the try/catch/finally statement.
type TryStmt struct {
	PosImpl
	Try     []Stmt
	Var     string
	Catch   []Stmt
	Finally []Stmt
}

// Execute the statement.
func (stmt *TryStmt) Execute(env *Env) (reflect.Value, error) {
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
			err = NewError(stmt.Catch[0], e1)
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
			err = NewError(stmt.Finally[0], e2)
		}
	}
	return NilValue, NewError(stmt, err)
}
