package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// LetsStmt defines a statement which defines multiple variables.
type LetsStmt struct {
	goblin.PosImpl
	Left     []goblin.Expr
	Operator string
	Right    []goblin.Expr
}

// Execute the statement.
func (stmt *LetsStmt) Execute(env *goblin.Env) (reflect.Value, error) {
	rv := goblin.NilValue
	var err error
	vs := make([]interface{}, 0, len(stmt.Right))
	for _, right := range stmt.Right {
		rv, err = right.Invoke(env)
		if err != nil {
			return rv, goblin.NewError(right, err)
		}
		if rv == goblin.NilValue {
			vs = append(vs, nil)
		} else if rv.IsValid() && rv.CanInterface() {
			vs = append(vs, rv.Interface())
		} else {
			vs = append(vs, nil)
		}
	}
	rvs := reflect.ValueOf(vs)
	if len(stmt.Left) > 1 && rvs.Len() == 1 {
		item := rvs.Index(0)
		if item.Kind() == reflect.Interface {
			item = item.Elem()
		}
		if item.Kind() == reflect.Slice {
			rvs = item
		}
	}
	for i, left := range stmt.Left {
		if i >= rvs.Len() {
			break
		}
		v := rvs.Index(i)
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		_, err = left.Assign(v, env)
		if err != nil {
			return rvs, goblin.NewError(left, err)
		}
	}
	if rvs.Len() == 1 {
		return rvs.Index(0), nil
	}
	return rvs, nil
}
