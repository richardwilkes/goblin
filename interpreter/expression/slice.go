package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Slice defines an array slice expression.
type Slice struct {
	interpreter.PosImpl
	Value interpreter.Expr
	Begin interpreter.Expr
	End   interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Slice) Invoke(env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
	}
	rb, err := expr.Begin.Invoke(env)
	if err != nil {
		return rb, interpreter.NewError(expr, err)
	}
	re, err := expr.End.Invoke(env)
	if err != nil {
		return re, interpreter.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(rb.Int())
		if ii < 0 || ii > v.Len() {
			return interpreter.NilValue, nil
		}
		ij := int(re.Int())
		if ij < 0 || ij > v.Len() {
			return v, nil
		}
		return v.Slice(ii, ij), nil
	}
	if v.Kind() == reflect.String {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		r := []rune(v.String())
		ii := int(rb.Int())
		if ii < 0 || ii >= len(r) {
			return interpreter.NilValue, nil
		}
		ij := int(re.Int())
		if ij < 0 || ij >= len(r) {
			return interpreter.NilValue, nil
		}
		return reflect.ValueOf(string(r[ii:ij])), nil
	}
	return v, interpreter.NewInvalidOperationError(expr)
}

// Assign a value to the expression and return it.
func (expr *Slice) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	v, err := expr.Value.Invoke(env)
	if err != nil {
		return v, interpreter.NewError(expr, err)
	}
	rb, err := expr.Begin.Invoke(env)
	if err != nil {
		return rb, interpreter.NewError(expr, err)
	}
	re, err := expr.End.Invoke(env)
	if err != nil {
		return re, interpreter.NewError(expr, err)
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return interpreter.NilValue, interpreter.NewArrayIndexShouldBeIntError(expr)
		}
		ii := int(rb.Int())
		if ii < 0 || ii >= v.Len() {
			return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
		}
		ij := int(re.Int())
		if ij < 0 || ij >= v.Len() {
			return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
		}
		vv := v.Slice(ii, ij)
		if !vv.CanSet() {
			return interpreter.NilValue, interpreter.NewCannotAssignError(expr)
		}
		vv.Set(rv)
		return rv, nil
	}
	return v, interpreter.NewInvalidOperationError(expr)
}
