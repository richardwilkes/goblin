package interpreter

import (
	"reflect"
)

// Stmt defines the required methods of a statement.
type Stmt interface {
	Pos
	// Execute the statement.
	Execute(env *Env) (reflect.Value, error)
}
