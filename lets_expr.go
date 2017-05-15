package goblin

import "reflect"

// LetsExpr defines an expresion to define multiple variables.
type LetsExpr struct {
	PosImpl
	LHSS     []Expr
	Operator string
	RHSS     []Expr
}

// Invoke the expression and return a result.
func (expr *LetsExpr) Invoke(env *Env) (reflect.Value, error) {
	rv := NilValue
	var err error
	vs := []interface{}{}
	for _, RHS := range expr.RHSS {
		rv, err = RHS.Invoke(env)
		if err != nil {
			return rv, NewError(RHS, err)
		}
		if rv == NilValue {
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
			return rvs, NewError(LHS, err)
		}
	}
	return rvs, nil
}

// Assign a value to the expression and return it.
func (expr *LetsExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
