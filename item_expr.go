package goblin

import "reflect"

// ItemExpr defines an expression that refers to a map or array item.
type ItemExpr struct {
	PosImpl
	Value Expr
	Index Expr
}

// Invoke the expression and return a result.
func (expr *ItemExpr) Invoke(env *Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	i, err := expr.Index.Invoke(env)
	if err != nil {
		return i, NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return NilValue, newArrayIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return NilValue, nil
		}
		return v.Index(ii), nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return NilValue, newMapKeyShouldBeStringError(expr)
		}
		return v.MapIndex(i), nil
	}
	if v.Kind() == reflect.String {
		rs := []rune(v.Interface().(string))
		ii := int(i.Int())
		if ii < 0 || ii >= len(rs) {
			return NilValue, nil
		}
		return reflect.ValueOf(rs[ii]), nil
	}
	return v, newInvalidOperationError(expr)
}

// Assign a value to the expression and return it.
func (expr *ItemExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	i, err := expr.Index.Invoke(env)
	if err != nil {
		return i, NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
			return NilValue, newArrayIndexShouldBeIntError(expr)
		}
		ii := int(i.Int())
		if ii < 0 || ii >= v.Len() {
			return NilValue, newCannotAssignError(expr)
		}
		vv := v.Index(ii)
		if !vv.CanSet() {
			return NilValue, newCannotAssignError(expr)
		}
		vv.Set(rv)
		return rv, nil
	}
	if v.Kind() == reflect.Map {
		if i.Kind() != reflect.String {
			return NilValue, newMapKeyShouldBeStringError(expr)
		}
		v.SetMapIndex(i, rv)
		return rv, nil
	}
	return v, newInvalidOperationError(expr)
}
