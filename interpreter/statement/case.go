package statement

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/interpreter"
)

// Case defines a case statement.
type Case struct {
	interpreter.PosImpl
	Expr  interpreter.Expr
	Stmts []interpreter.Stmt
}

func (stmt *Case) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "case %v:", stmt.Expr)
	for _, stmt := range stmt.Stmts {
		fmt.Fprintf(&buffer, "\n        %v", stmt)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Case) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewCannotExecuteError(stmt)
}
