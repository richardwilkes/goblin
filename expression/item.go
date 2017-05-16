package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Item defines an expression that refers to a map or array item.
type Item struct {
	goblin.PosImpl
	Value goblin.Expr
	Index goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Item) Invoke(env *goblin.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, goblin.NewError(expr, err)
	}
	i, err := expr.Index.Invoke(env)
	if err != nil {
		return i, goblin.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return goblin.NilValue, nil
		}
		return v.Index(ii), nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return goblin.NilValue, goblin.NewMapKeyShouldBeStringError(expr)
		}
		return v.MapIndex(i), nil
	}
	if v.Kind() == reflect.String {
		rs := []rune(v.Interface().(string))
		ii := int(i.Int())
		if ii < 0 || ii >= len(rs) {
			return goblin.NilValue, nil
		}
		return reflect.ValueOf(rs[ii]), nil
	}
	return v, goblin.NewInvalidOperationError(expr)
}

// Assign a value to the expression and return it.
func (expr *Item) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, goblin.NewError(expr, err)
	}
	i, err := expr.Index.Invoke(env)
	if err != nil {
		return i, goblin.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return goblin.NilValue, goblin.NewCannotAssignError(expr)
		}
		vv := v.Index(ii)
		if !vv.CanSet() {
			return goblin.NilValue, goblin.NewCannotAssignError(expr)
		}
		vv.Set(rv)
		return rv, nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return goblin.NilValue, goblin.NewMapKeyShouldBeStringError(expr)
		}
		v.SetMapIndex(i, rv)
		return rv, nil
	}
	return v, goblin.NewInvalidOperationError(expr)
}
