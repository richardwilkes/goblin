package statement

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Case defines a case statement.
type Case struct {
	ast.PosImpl
	Expr  ast.Expr
	Stmts []ast.Stmt
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
func (stmt *Case) Execute(scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewCannotExecuteError(stmt)
}
