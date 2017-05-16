package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Member defines a member reference expression.
type Member struct {
	interpreter.PosImpl
	Expr interpreter.Expr
	Name string
}

// Invoke the expression and return a result.
func (expr *Member) Invoke(env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if v.IsValid() && v.CanInterface() {
		if vme, ok := v.Interface().(*interpreter.Env); ok {
			m, err := vme.Get(expr.Name)
			if !m.IsValid() || err != nil {
				return interpreter.NilValue, interpreter.NewNamedInvalidOperationError(expr, expr.Name)
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
				return interpreter.NilValue, interpreter.NewNamedInvalidOperationError(expr, expr.Name)
			}
		} else if kind == reflect.Map {
			m = v.MapIndex(reflect.ValueOf(expr.Name))
			if !m.IsValid() {
				return interpreter.NilValue, interpreter.NewNamedInvalidOperationError(expr, expr.Name)
			}
		} else {
			return interpreter.NilValue, interpreter.NewNamedInvalidOperationError(expr, expr.Name)
		}
	}
	return m, nil
}

// Assign a value to the expression and return it.
func (expr *Member) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.Expr.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Slice {
		v = v.Index(0)
	}
	if !v.IsValid() {
		return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		v = v.FieldByName(expr.Name)
		if !v.CanSet() {
			return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
		}
		v.Set(rv)
	} else if v.Kind() == reflect.Map {
		v.SetMapIndex(reflect.ValueOf(expr.Name), rv)
	} else {
		if !v.CanSet() {
			return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
		}
		v.Set(rv)
	}
	return v, nil
}
