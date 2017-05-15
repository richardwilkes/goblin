package goblin

import "reflect"

// Expr defines the required methods of an expression.
type Expr interface {
	Pos
	// Invoke the expression and return a result.
	Invoke(env *Env) (reflect.Value, error)
	// Assign a value to the expression and return it.
	Assign(rv reflect.Value, env *Env) (reflect.Value, error)
}

// Type defines a type.
type Type struct {
	Name string
}
