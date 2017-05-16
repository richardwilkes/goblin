package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/util"
)

// Return defines the return statement.
type Return struct {
	goblin.PosImpl
	Exprs []goblin.Expr
}

// Execute the statement.
func (stmt *Return) Execute(env *goblin.Env) (reflect.Value, error) {
	rvs := []interface{}{}
	switch len(stmt.Exprs) {
	case 0:
		return goblin.NilValue, goblin.ErrReturn
	case 1:
		rv, err := stmt.Exprs[0].Invoke(env)
		if err != nil {
			return rv, goblin.NewError(stmt, err)
		}
		return rv, goblin.ErrReturn
	}
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(env)
		if err != nil {
			return rv, goblin.NewError(stmt, err)
		}
		if util.IsNil(rv) {
			rvs = append(rvs, nil)
		} else if rv.IsValid() {
			rvs = append(rvs, rv.Interface())
		} else {
			rvs = append(rvs, nil)
		}
	}
	return reflect.ValueOf(rvs), goblin.ErrReturn
}
