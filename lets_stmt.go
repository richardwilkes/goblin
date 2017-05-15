package goblin

import "reflect"

// LetsStmt defines a statement which defines multiple variables.
type LetsStmt struct {
	PosImpl
	LHSS     []Expr
	Operator string
	RHSS     []Expr
}

// Execute the statement.
func (stmt *LetsStmt) Execute(env *Env) (reflect.Value, error) {
	rv := NilValue
	var err error
	vs := make([]interface{}, 0, len(stmt.RHSS))
	for _, RHS := range stmt.RHSS {
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
	if len(stmt.LHSS) > 1 && rvs.Len() == 1 {
		item := rvs.Index(0)
		if item.Kind() == reflect.Interface {
			item = item.Elem()
		}
		if item.Kind() == reflect.Slice {
			rvs = item
		}
	}
	for i, LHS := range stmt.LHSS {
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
	if rvs.Len() == 1 {
		return rvs.Index(0), nil
	}
	return rvs, nil
}
