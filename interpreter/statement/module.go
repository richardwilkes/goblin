package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Module defines a module statement.
type Module struct {
	interpreter.PosImpl
	Name  string
	Stmts []interpreter.Stmt
}

// Execute the statement.
func (stmt *Module) Execute(env *interpreter.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	newEnv.SetName(stmt.Name)
	rv, err := newEnv.Run(stmt.Stmts)
	if err != nil {
		return rv, interpreter.NewError(stmt, err)
	}
	env.DefineGlobal(stmt.Name, reflect.ValueOf(newEnv))
	return rv, nil
}
