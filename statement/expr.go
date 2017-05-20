package statement

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Expression defines an expression statement.
type Expression struct {
	ast.PosImpl
	Expr ast.Expr
}

func (stmt *Expression) String() string {
	return fmt.Sprint(stmt.Expr)
}

// Execute the statement.
func (stmt *Expression) Execute(scope ast.Scope) (reflect.Value, error) {
	rv, err := stmt.Expr.Invoke(scope)
	if err != nil {
		return rv, ast.NewError(stmt, err)
	}
	return rv, nil
}
