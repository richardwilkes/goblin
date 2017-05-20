package statement

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Throw defines the throw statement.
type Throw struct {
	ast.PosImpl
	Expr ast.Expr
}

func (stmt *Throw) String() string {
	return fmt.Sprintf("throw %v", stmt.Expr)
}

// Execute the statement.
func (stmt *Throw) Execute(scope ast.Scope) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(stmt, err)
	}
	if !rv.IsValid() {
		return ast.NilValue, ast.NewError(stmt, err)
	}
	return rv, ast.NewStringError(stmt, fmt.Sprint(rv.Interface()))
}
