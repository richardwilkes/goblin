package ast

import (
	"reflect"
)

// Stmt defines the required methods of a statement.
type Stmt interface {
	Pos
	// Execute the statement.
	Execute(scope Scope) (reflect.Value, error)
}
