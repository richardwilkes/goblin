package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Module defines a module statement.
type Module struct {
	interpreter.PosImpl
	Name  string
	Stmts []interpreter.Stmt
}

func (stmt *Module) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "module %s {", stmt.Name)
	for _, stmt := range stmt.Stmts {
		fmt.Fprintf(&buffer, "\n    %v", stmt)
	}
	buffer.WriteString("\n}")
	return buffer.String()
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
