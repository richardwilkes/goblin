package goblin

import "reflect"

// VarStmt defines a variable definition statement.
type VarStmt struct {
	PosImpl
	Names []string
	Exprs []Expr
}

// Execute the statement.
func (stmt *VarStmt) Execute(env *Env) (reflect.Value, error) {
	rvs := make([]reflect.Value, 0, len(stmt.Exprs))
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(env)
		if err != nil {
			return rv, NewError(expr, err)
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
