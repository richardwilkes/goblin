package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Default defines the default case statement.
type Default struct {
	interpreter.PosImpl
	Stmts []interpreter.Stmt
}

func (stmt *Default) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("default:")
	for _, stmt := range stmt.Stmts {
		fmt.Fprintf(&buffer, "\n        %v", stmt)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Default) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewCannotExecuteError(stmt)
}
