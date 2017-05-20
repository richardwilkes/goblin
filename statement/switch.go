package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Switch defines a switch statement.
type Switch struct {
	ast.PosImpl
	Expr  ast.Expr
	Cases []ast.Stmt
}

func (stmt *Switch) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "switch %v {", stmt.Expr)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Cases {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *Switch) Execute(scope ast.Scope) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(stmt, err)
	}
	done := false
	var defaultStmt *Default
	for _, ss := range stmt.Cases {
		if ssd, ok := ss.(*Default); ok {
			defaultStmt = ssd
			continue
		}
		caseStmt, ok := ss.(*Case)
		if !ok {
			return ast.NilValue, ast.NewError(stmt, ast.ErrBadSyntax)
		}
		cv, lerr := caseStmt.Expr.Invoke(scope)
		if lerr != nil {
			return rv, ast.NewError(stmt, lerr)
		}
		if !util.Equal(rv, cv) {
			continue
		}
		rv, err = scope.Run(caseStmt.Stmts)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
		done = true
		break
	}
	if !done && defaultStmt != nil {
		rv, err = scope.Run(defaultStmt.Stmts)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
	}
	return rv, nil
}
