package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// Variable defines a variable definition statement.
type Variable struct {
	goblin.PosImpl
	Names []string
	Exprs []goblin.Expr
}

// Execute the statement.
func (stmt *Variable) Execute(env *goblin.Env) (reflect.Value, error) {
	rvs := make([]reflect.Value, 0, len(stmt.Exprs))
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(env)
		if err != nil {
			return rv, goblin.NewError(expr, err)
		}
		rvs = append(rvs, rv)
	}
	result := make([]interface{}, 0, len(rvs))
	for i, name := range stmt.Names {
		if i < len(rvs) {
			env.Define(name, rvs[i])
			result = append(result, rvs[i].Interface())
		}
	}
	return reflect.ValueOf(result), nil
}
