package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
)

// Variables defines a statement which defines multiple variables.
type Variables struct {
	interpreter.PosImpl
	Left     []interpreter.Expr
	Operator string
	Right    []interpreter.Expr
}

// Execute the statement.
func (stmt *Variables) Execute(env *interpreter.Env) (reflect.Value, error) {
	rv := interpreter.NilValue
	var err error
	vs := make([]interface{}, 0, len(stmt.Right))
	for _, right := range stmt.Right {
		rv, err = right.Invoke(env)
		if err != nil {
			return rv, interpreter.NewError(right, err)
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
			return rvs, interpreter.NewError(left, err)
		}
	}
	if rvs.Len() == 1 {
		return rvs.Index(0), nil
	}
	return rvs, nil
}
