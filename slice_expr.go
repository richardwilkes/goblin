package goblin

import "reflect"

// SliceExpr defines an array slice expression.
type SliceExpr struct {
	PosImpl
	Value Expr
	Begin Expr
	End   Expr
}

// Invoke the expression and return a result.
func (expr *SliceExpr) Invoke(env *Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	rb, err := expr.Begin.Invoke(env)
	if err != nil {
		return rb, NewError(expr, err)
	}
	re, err := expr.End.Invoke(env)
	if err != nil {
		return re, NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return NilValue, NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return NilValue, NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(rb.Int())
		if ii < 0 || ii > v.Len() {
			return NilValue, nil
		}
		ij := int(re.Int())
		if ij < 0 || ij > v.Len() {
			return v, nil
		}
		return v.Slice(ii, ij), nil
	}
	if v.Kind() == reflect.String {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return NilValue, NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return NilValue, NewArrayIndexShouldBeIntError(expr)
		}
		r := []rune(v.String())
		ii := int(rb.Int())
		if ii < 0 || ii >= len(r) {
			return NilValue, nil
		}
		ij := int(re.Int())
		if ij < 0 || ij >= len(r) {
			return NilValue, nil
		}
		return reflect.ValueOf(string(r[ii:ij])), nil
	}
	return v, NewInvalidOperationError(expr)
}

// Assign a value to the expression and return it.
func (expr *SliceExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, NewError(expr, err)
	}
	rb, err := expr.Begin.Invoke(env)
	if err != nil {
		return rb, NewError(expr, err)
	}
	re, err := expr.End.Invoke(env)
	if err != nil {
		return re, NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return NilValue, NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return NilValue, NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(rb.Int())
		if ii < 0 || ii >= v.Len() {
			return NilValue, NewCannotAssignError(expr)
		}
		ij := int(re.Int())
		if ij < 0 || ij >= v.Len() {
			return NilValue, NewCannotAssignError(expr)
		}
		vv := v.Slice(ii, ij)
		if !vv.CanSet() {
			return NilValue, NewCannotAssignError(expr)
		}
		vv.Set(rv)
		return rv, nil
	}
	return v, NewInvalidOperationError(expr)
}
