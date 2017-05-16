package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin"
)

// VarStmt defines a variable definition statement.
type VarStmt struct {
	goblin.PosImpl
	Names []string
	Exprs []goblin.Expr
}

// Execute the statement.
func (stmt *VarStmt) Execute(env *goblin.Env) (reflect.Value, error) {
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
