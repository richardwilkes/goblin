package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Member defines a member reference expression.
type Member struct {
	ast.PosImpl
	Expr ast.Expr
	Name string
}

func (expr *Member) String() string {
	return fmt.Sprintf("%v.%s", expr.Expr, expr.Name)
}

// Invoke the expression and return a result.
func (expr *Member) Invoke(scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if v.IsValid() && v.CanInterface() {
		if vme, ok := v.Interface().(ast.Scope); ok {
			m, err := vme.Get(expr.Name)
			if !m.IsValid() || err != nil {
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
			}
			return m, nil
		}
	}

	m := v.MethodByName(expr.Name)
	if !m.IsValid() {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		kind := v.Kind()
		if kind == reflect.Struct {
			m = v.FieldByName(expr.Name)
			if !m.IsValid() {
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
			}
		} else if kind == reflect.Map {
			m = v.MapIndex(reflect.ValueOf(expr.Name))
			if !m.IsValid() {
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
			}
		} else {
			return ast.NilValue, ast.NewNamedInvalidOperationError(expr, expr.Name)
		}
	}
	return m, nil
}

// Assign a value to the expression and return it.
func (expr *Member) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(scope)
	if err != nil {
		return v, ast.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if !v.IsValid() {
		return ast.NilValue, ast.NewCannotAssignError(expr)
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		v = v.FieldByName(expr.Name)
		if !v.CanSet() {
			return ast.NilValue, ast.NewCannotAssignError(expr)
		}
		v.Set(rv)
	} else if v.Kind() == reflect.Map {
		v.SetMapIndex(reflect.ValueOf(expr.Name), rv)
	} else {
		if !v.CanSet() {
			return ast.NilValue, ast.NewCannotAssignError(expr)
		}
		v.Set(rv)
	}
	return v, nil
}