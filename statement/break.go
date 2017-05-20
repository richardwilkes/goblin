package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Break defines a break statement.
type Break struct {
	ast.PosImpl
}

func (stmt *Break) String() string {
	return "break"
}

// Execute the statement.
func (stmt *Break) Execute(scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.ErrBreak
}
