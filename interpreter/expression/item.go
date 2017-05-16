package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Item defines an expression that refers to a map or array item.
type Item struct {
	interpreter.PosImpl
	Value interpreter.Expr
	Index interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Item) Invoke(env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
	}
	i, err := expr.Index.Invoke(env)
	if err != nil {
		return i, interpreter.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return interpreter.NilValue, nil
		}
		return v.Index(ii), nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return interpreter.NilValue, interpreter.NewMapKeyShouldBeStringError(expr)
		}
		return v.MapIndex(i), nil
	}
	if v.Kind() == reflect.String {
		rs := []rune(v.Interface().(string))
		ii := int(i.Int())
		if ii < 0 || ii >= len(rs) {
			return interpreter.NilValue, nil
		}
		return reflect.ValueOf(rs[ii]), nil
	}
	return v, interpreter.NewInvalidOperationError(expr)
}

// Assign a value to the expression and return it.
func (expr *Item) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
	}
	i, err := expr.Index.Invoke(env)
	if err != nil {
		return i, interpreter.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
		}
		vv := v.Index(ii)
		if !vv.CanSet() {
			return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
		}
		vv.Set(rv)
		return rv, nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return interpreter.NilValue, interpreter.NewMapKeyShouldBeStringError(expr)
		}
		v.SetMapIndex(i, rv)
		return rv, nil
	}
	return v, interpreter.NewInvalidOperationError(expr)
}
