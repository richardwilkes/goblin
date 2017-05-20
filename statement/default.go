package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Default defines the default case statement.
type Default struct {
	ast.PosImpl
	Stmts []ast.Stmt
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
func (stmt *Default) Execute(scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewCannotExecuteError(stmt)
}
