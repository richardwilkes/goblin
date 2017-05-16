package goblin

import (
	"reflect"
)

// AddrExpr defines a referencing address expression.
type AddrExpr struct {
	PosImpl
	Expr Expr
}

// Invoke the expression and return a result.
func (expr *AddrExpr) Invoke(env *Env) (reflect.Value, error) {
	v := NilValue
	var err error
	switch ee := expr.Expr.(type) {
	case *IdentExpr:
		v, err = env.Get(ee.Lit)
		if err != nil {
			return v, err
		}
	case *MemberExpr:
		v, err = ee.Expr.Invoke(env)
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
				m, err := vme.Get(ee.Name)
				if !m.IsValid() || err != nil {
					return NilValue, NewNamedInvalidOperationError(expr, ee.Name)
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
					return NilValue, NewNamedInvalidOperationError(expr, ee.Name)
				}
			} else if kind == reflect.Map {
				m = v.MapIndex(reflect.ValueOf(ee.Name))
				if !m.IsValid() {
					return NilValue, NewNamedInvalidOperationError(expr, ee.Name)
				}
			} else {
				return NilValue, NewNamedInvalidOperationError(expr, ee.Name)
			}
			v = m
		} else {
			v = m
		}
	default:
		return NilValue, NewStringError(expr, "Invalid operation for the value")
	}
	if !v.CanAddr() {
		i := v.Interface()
		return reflect.ValueOf(&i), nil
	}
	return v.Addr(), nil
}

// Assign a value to the expression and return it.
func (expr *AddrExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, NewInvalidOperationError(expr)
}
