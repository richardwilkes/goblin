package expression

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Vars defines an expression that defines multiple variables.
type Vars struct {
	interpreter.PosImpl
	LHSS     []interpreter.Expr
	Operator string
	RHSS     []interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *Vars) Invoke(env *interpreter.Env) (reflect.Value, error) {
	rv := interpreter.NilValue
	var err error
	vs := []interface{}{}
	for _, RHS := range expr.RHSS {
		rv, err = RHS.Invoke(env)
		if err != nil {
			return rv, interpreter.NewError(RHS, err)
		}
		if rv == interpreter.NilValue {
			vs = append(vs, nil)
		} else if rv.IsValid() && rv.CanInterface() {
			vs = append(vs, rv.Interface())
		} else {
			vs = append(vs, nil)
		}
	}
	rvs := reflect.ValueOf(vs)
	for i, LHS := range expr.LHSS {
		if i >= rvs.Len() {
			break
		}
		v := rvs.Index(i)
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		_, err = LHS.Assign(v, env)
		if err != nil {
			return rvs, interpreter.NewError(LHS, err)
		}
	}
	return rvs, nil
}

// Assign a value to the expression and return it.
func (expr *Vars) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
