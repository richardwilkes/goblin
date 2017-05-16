package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// ModuleStmt defines a module statement.
type ModuleStmt struct {
	goblin.PosImpl
	Name  string
	Stmts []goblin.Stmt
}

// Execute the statement.
func (stmt *ModuleStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	newEnv.SetName(stmt.Name)
	rv, err := newEnv.Run(stmt.Stmts)
	if err != nil {
		return rv, goblin.NewError(stmt, err)
	}
	env.DefineGlobal(stmt.Name, reflect.ValueOf(newEnv))
	return rv, nil
}
