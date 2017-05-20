package expression

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Deref defines a dereferencing address expression.
type Deref struct {
	ast.PosImpl
	Expr ast.Expr
}

func (expr *Deref) String() string {
	return fmt.Sprintf("*%v", expr.Expr)
}

// Invoke the expression and return a result.
func (expr *Deref) Invoke(scope ast.Scope) (reflect.Value, error) {
	v := ast.NilValue
	var err error
	switch ee := expr.Expr.(type) {
	case *Ident:
		v, err = scope.Get(ee.Lit)
		if err != nil {
			return v, err
		}
	case *Member:
		v, err = ee.Expr.Invoke(scope)
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
				m, err := vme.Get(ee.Name)
				if !m.IsValid() || err != nil {
					return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
				}
				return m, nil
			}
		}

		m := v.MethodByName(ee.Name)
		if !m.IsValid() {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			kind := v.Kind()
			if kind == reflect.Struct {
				m = v.FieldByName(ee.Name)
				if !m.IsValid() {
					return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
				}
			} else if kind == reflect.Map {
				m = v.MapIndex(reflect.ValueOf(ee.Name))
				if !m.IsValid() {
					return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
				}
			} else {
				return ast.NilValue, ast.NewNamedInvalidOperationError(expr, ee.Name)
			}
			v = m
		} else {
			v = m
		}
	default:
		return ast.NilValue, ast.NewStringError(expr, "Invalid operation for the value")
	}
	if v.Kind() != reflect.Ptr {
		return ast.NilValue, ast.NewStringError(expr, "Cannot deference for the value")
	}
	return v.Elem(), nil
}

// Assign a value to the expression and return it.
func (expr *Deref) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}