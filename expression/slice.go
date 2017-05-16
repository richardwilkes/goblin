package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Slice defines an array slice expression.
type Slice struct {
	goblin.PosImpl
	Value goblin.Expr
	Begin goblin.Expr
	End   goblin.Expr
}

// Invoke the expression and return a result.
func (expr *Slice) Invoke(env *goblin.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, goblin.NewError(expr, err)
	}
	rb, err := expr.Begin.Invoke(env)
	if err != nil {
		return rb, goblin.NewError(expr, err)
	}
	re, err := expr.End.Invoke(env)
	if err != nil {
		return re, goblin.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(rb.Int())
		if ii < 0 || ii > v.Len() {
			return goblin.NilValue, nil
		}
		ij := int(re.Int())
		if ij < 0 || ij > v.Len() {
			return v, nil
		}
		return v.Slice(ii, ij), nil
	}
	if v.Kind() == reflect.String {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		r := []rune(v.String())
		ii := int(rb.Int())
		if ii < 0 || ii >= len(r) {
			return goblin.NilValue, nil
		}
		ij := int(re.Int())
		if ij < 0 || ij >= len(r) {
			return goblin.NilValue, nil
		}
		return reflect.ValueOf(string(r[ii:ij])), nil
	}
	return v, goblin.NewInvalidOperationError(expr)
}

// Assign a value to the expression and return it.
func (expr *Slice) Assign(rv reflect.Value, env *goblin.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, goblin.NewError(expr, err)
	}
	rb, err := expr.Begin.Invoke(env)
	if err != nil {
		return rb, goblin.NewError(expr, err)
	}
	re, err := expr.End.Invoke(env)
	if err != nil {
		return re, goblin.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return goblin.NilValue, goblin.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(rb.Int())
		if ii < 0 || ii >= v.Len() {
			return goblin.NilValue, goblin.NewCannotAssignError(expr)
		}
		ij := int(re.Int())
		if ij < 0 || ij >= v.Len() {
			return goblin.NilValue, goblin.NewCannotAssignError(expr)
		}
		vv := v.Slice(ii, ij)
		if !vv.CanSet() {
			return goblin.NilValue, goblin.NewCannotAssignError(expr)
		}
		vv.Set(rv)
		return rv, nil
	}
	return v, goblin.NewInvalidOperationError(expr)
}
