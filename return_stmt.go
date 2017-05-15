package goblin

import "reflect"

// ReturnStmt defines the return statement.
type ReturnStmt struct {
	PosImpl
	Exprs []Expr
}

// Execute the statement.
func (stmt *ReturnStmt) Execute(env *Env) (reflect.Value, error) {
	rvs := []interface{}{}
	switch len(stmt.Exprs) {
	case 0:
		return NilValue, ErrReturn
	case 1:
		rv, err := stmt.Exprs[0].Invoke(env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		return rv, ErrReturn
	}
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if isNil(rv) {
			rvs = append(rvs, nil)
		} else if rv.IsValid() {
			rvs = append(rvs, rv.Interface())
		} else {
			rvs = append(rvs, nil)
		}
	}
	return reflect.ValueOf(rvs), ErrReturn
}
