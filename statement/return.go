package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Return defines the return statement.
type Return struct {
	ast.PosImpl
	Exprs []ast.Expr
}

func (stmt *Return) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("return")
	for i, one := range stmt.Exprs {
		if i != 0 {
			buffer.WriteString(",")
		}
		fmt.Fprintf(&buffer, " %v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Return) Execute(scope ast.Scope) (reflect.Value, error) {
	rvs := []interface{}{}
	switch len(stmt.Exprs) {
	case 0:
		return ast.NilValue, ast.ErrReturn
	case 1:
		rv, err := stmt.Exprs[0].Invoke(scope)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
		return rv, ast.ErrReturn
	}
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(scope)
		if err != nil {
			return rv, ast.NewError(stmt, err)
		}
		switch {
		case util.IsNil(rv):
			rvs = append(rvs, nil)
		case rv.IsValid():
			rvs = append(rvs, rv.Interface())
		default:
			rvs = append(rvs, nil)
		}
	}
	return reflect.ValueOf(rvs), ast.ErrReturn
}
