package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Assoc defines an operator association expression.
type Assoc struct {
	ast.PosImpl
	Left     ast.Expr
	Operator string
	Right    ast.Expr
}

func (expr *Assoc) String() string {
	switch expr.Operator {
	case "++", "--":
		return fmt.Sprintf("%v%s", expr.Left, expr.Operator)
	default:
		return fmt.Sprintf("%v %s %v", expr.Left, expr.Operator, expr.Right)
	}
}

// Invoke the expression and return a result.
func (expr *Assoc) Invoke(scope ast.Scope) (reflect.Value, error) {
	switch expr.Operator {
	case "++":
		if aLHS, ok := expr.Left.(*Ident); ok {
			v, err := scope.Get(aLHS.Lit)
			if err != nil {
				return v, err
			}
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(util.ToFloat64(v) + 1.0)
			} else {
				v = reflect.ValueOf(util.ToInt64(v) + 1)
			}
			if scope.Set(aLHS.Lit, v) != nil {
				scope.Define(aLHS.Lit, v)
			}
			return v, nil
		}
	case "--":
		if aLHS, ok := expr.Left.(*Ident); ok {
			v, err := scope.Get(aLHS.Lit)
			if err != nil {
				return v, err
			}
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(util.ToFloat64(v) - 1.0)
			} else {
				v = reflect.ValueOf(util.ToInt64(v) - 1)
			}
			if scope.Set(aLHS.Lit, v) != nil {
				scope.Define(aLHS.Lit, v)
			}
			return v, nil
		}
	}

	binop := &BinOp{Left: expr.Left, Operator: expr.Operator[0:1], Right: expr.Right}
	v, err := binop.Invoke(scope)
	if err != nil {
		return v, err
	}

	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return expr.Left.Assign(v, scope)
}

// Assign a value to the expression and return it.
func (expr *Assoc) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
