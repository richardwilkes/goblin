package statement

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
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
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, stmt := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", stmt)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Case) Execute(env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewCannotExecuteError(stmt)
}
