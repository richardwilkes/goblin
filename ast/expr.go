package ast

import "reflect"

// Expr defines the required methods of an expression.
type Expr interface {
	Pos
	// Invoke the expression and return a result.
	Invoke(scope Scope) (reflect.Value, error)
	// Assign a value to the expression and return it.
	Assign(rv reflect.Value, scope Scope) (reflect.Value, error)
}

// Type defines a type.
type Type struct {
	Name string
}
