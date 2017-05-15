package goblin

import (
	"reflect"
)

// MemberExpr defines a member reference expression.
type MemberExpr struct {
	PosImpl
	Expr Expr
	Name string
}

// Invoke the expression and return a result.
func (expr *MemberExpr) Invoke(env *Env) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if v.IsValid() && v.CanInterface() {
		if vme, ok := v.Interface().(*Env); ok {
			m, err := vme.Get(expr.Name)
			if !m.IsValid() || err != nil {
				return NilValue, newNamedInvalidOperationError(expr, expr.Name)
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
				return NilValue, newNamedInvalidOperationError(expr, expr.Name)
			}
		} else if kind == reflect.Map {
			m = v.MapIndex(reflect.ValueOf(expr.Name))
			if !m.IsValid() {
				return NilValue, newNamedInvalidOperationError(expr, expr.Name)
			}
		} else {
			return NilValue, newNamedInvalidOperationError(expr, expr.Name)
		}
	}
	return m, nil
}

// Assign a value to the expression and return it.
func (expr *MemberExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if !v.IsValid() {
		return NilValue, newCannotAssignError(expr)
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		v = v.FieldByName(expr.Name)
		if !v.CanSet() {
			return NilValue, newCannotAssignError(expr)
		}
		v.Set(rv)
	} else if v.Kind() == reflect.Map {
		v.SetMapIndex(reflect.ValueOf(expr.Name), rv)
	} else {
		if !v.CanSet() {
			return NilValue, newCannotAssignError(expr)
		}
		v.Set(rv)
	}
	return v, nil
}
