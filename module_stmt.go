package goblin

import "reflect"

// ModuleStmt defines a module statement.
type ModuleStmt struct {
	PosImpl
	Name  string
	Stmts []Stmt
}

// Execute the statement.
func (stmt *ModuleStmt) Execute(env *Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	newEnv.SetName(stmt.Name)
	rv, err := newEnv.Run(stmt.Stmts)
	if err != nil {
		return rv, NewError(stmt, err)
	}
	env.DefineGlobal(stmt.Name, reflect.ValueOf(newEnv))
	return rv, nil
}
