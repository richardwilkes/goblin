package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Loop defines a loop statement.
type Loop struct {
	ast.PosImpl
	Expr  ast.Expr
	Stmts []ast.Stmt
}

func (stmt *Loop) String() string {
	var buffer bytes.Buffer
	if stmt.Expr != nil {
		fmt.Fprintf(&buffer, "for %v {", stmt.Expr)
	} else {
		buffer.WriteString("for {")
	}
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *Loop) Execute(scope ast.Scope) (reflect.Value, error) {
	newScope := scope.NewScope()
	defer newScope.Destroy()
	for {
		if stmt.Expr != nil {
			ev, ee := stmt.Expr.Invoke(newScope)
			if ee != nil {
				return ev, ee
			}
			if !util.ToBool(ev) {
				break
			}
		}

		rv, err := newScope.Run(stmt.Stmts)
		if err != nil {
			if err == ast.ErrBreak {
				break
			}
			if err == ast.ErrContinue {
				continue
			}
			if err == ast.ErrReturn {
				return rv, err
			}
			return rv, ast.NewError(stmt, err)
		}
	}
	return ast.NilValue, nil
}
