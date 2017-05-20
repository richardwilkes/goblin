package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Continue defines the continue statement.
type Continue struct {
	ast.PosImpl
}

func (stmt *Continue) String() string {
	return "continue"
}

// Execute the statement.
func (stmt *Continue) Execute(scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.ErrContinue
}
