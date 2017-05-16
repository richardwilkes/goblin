package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Addr defines a referencing address expression.
type Addr struct {
	goblin.PosImpl
	Expr goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Addr) Invoke(env *goblin.Env) (reflect.Value, error) {
	v := goblin.NilValue
	var err error
	switch ee := expr.Expr.(type) {
	case *Ident:
		v, err = env.Get(ee.Lit)
		if err != nil {
			return v, err
		}
	case *Member:
		v, err = ee.Expr.Invoke(env)
		if err != nil {
			return v, goblin.NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Slice {
			v = v.Index(0)
		}
		if v.IsValid() && v.CanInterface() {
			if vme, ok := v.Interface().(*goblin.Env); ok {
				m, err := vme.Get(ee.Name)
				if !m.IsValid() || err != nil {
					return goblin.NilValue, goblin.NewNamedInvalidOperationError(expr, ee.Name)
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
					return goblin.NilValue, goblin.NewNamedInvalidOperationError(expr, ee.Name)
				}
			} else if kind == reflect.Map {
				m = v.MapIndex(reflect.ValueOf(ee.Name))
				if !m.IsValid() {
					return goblin.NilValue, goblin.NewNamedInvalidOperationError(expr, ee.Name)
				}
			} else {
				return goblin.NilValue, goblin.NewNamedInvalidOperationError(expr, ee.Name)
			}
			v = m
		} else {
			v = m
		}
	default:
		return goblin.NilValue, goblin.NewStringError(expr, "Invalid operation for the value")
	}
	if !v.CanAddr() {
		i := v.Interface()
		return reflect.ValueOf(&i), nil
	}
	return v.Addr(), nil
}

// Assign a value to the expression and return it.
func (expr *Addr) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	return goblin.NilValue, goblin.NewInvalidOperationError(expr)
}
