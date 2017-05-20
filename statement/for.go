package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// For defines a for statement.
type For struct {
	ast.PosImpl
	Var   string
	Value ast.Expr
	Stmts []ast.Stmt
}

func (stmt *For) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "for %s in %v {", stmt.Var, stmt.Value)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *For) Execute(scope ast.Scope) (reflect.Value, error) {
	val, ee := stmt.Value.Invoke(scope)
	if ee != nil {
		return val, ee
	}
	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		newScope := scope.NewScope()
		defer newScope.Destroy()

		for i := 0; i < val.Len(); i++ {
			iv := val.Index(i)
			if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
				iv = iv.Elem()
			}
			newScope.Define(stmt.Var, iv)
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
	return ast.NilValue, ast.NewStringError(stmt, "Invalid operation for non-array value")
}
