package statement

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// If defines an if/else statement.
type If struct {
	ast.PosImpl
	If     ast.Expr
	Then   []ast.Stmt
	ElseIf []ast.Stmt
	Else   []ast.Stmt
}

func (stmt *If) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "if %v {", stmt.If)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Then {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	if len(stmt.ElseIf) > 0 {
		for _, one := range stmt.ElseIf {
			fmt.Fprintf(&buffer, " else %v", one)
		}
	}
	if len(stmt.Else) > 0 {
		buffer.WriteString(" else {")
		for _, one := range stmt.Else {
			fmt.Fprintf(prefixer, "\n%v", one)
		}
		buffer.WriteString("\n}")
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *If) Execute(scope ast.Scope) (reflect.Value, error) {
	// If
	rv, err := stmt.If.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(stmt, err)
	}
	if util.ToBool(rv) {
		// Then
		newScope := scope.NewScope()
		defer newScope.Destroy()
		rv, err = newScope.Run(stmt.Then)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
		return rv, nil
	}
	done := false
	if len(stmt.ElseIf) > 0 {
		for _, stmt := range stmt.ElseIf {
			stmtIf, ok := stmt.(*If)
			if !ok {
				return ast.NilValue, ast.NewError(stmt, ast.ErrBadSyntax)
			}
			// Else If
			rv, err = stmtIf.If.Invoke(scope)
			if err != nil {
				return rv, ast.NewError(stmt, err)
			}
			if !util.ToBool(rv) {
				continue
			}
			// Else If Then
			done = true
			rv, err = scope.Run(stmtIf.Then)
			if err != nil {
				return rv, ast.NewError(stmt, err)
			}
			break
		}
	}
	if !done && len(stmt.Else) > 0 {
		// Else
		newScope := scope.NewScope()
		defer newScope.Destroy()
		rv, err = newScope.Run(stmt.Else)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
	}
	return rv, nil
}
