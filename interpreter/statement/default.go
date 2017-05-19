package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Default defines the default case statement.
type Default struct {
	interpreter.PosImpl
	Stmts []interpreter.Stmt
}

func (stmt *Default) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("default:")
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Default) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewCannotExecuteError(stmt)
}
