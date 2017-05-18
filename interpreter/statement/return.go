package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Return defines the return statement.
type Return struct {
	interpreter.PosImpl
	Exprs []interpreter.Expr
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
func (stmt *Return) Execute(env *interpreter.Env) (reflect.Value, error) {
	rvs := []interface{}{}
	switch len(stmt.Exprs) {
	case 0:
		return interpreter.NilValue, interpreter.ErrReturn
	case 1:
		rv, err := stmt.Exprs[0].Invoke(env)
		if err != nil {
			return rv, interpreter.NewError(stmt, err)
		}
		return rv, interpreter.ErrReturn
	}
	for _, expr := range stmt.Exprs {
		rv, err := expr.Invoke(env)
		if err != nil {
			return rv, interpreter.NewError(stmt, err)
		}
		if util.IsNil(rv) {
			rvs = append(rvs, nil)
		} else if rv.IsValid() {
			rvs = append(rvs, rv.Interface())
		} else {
			rvs = append(rvs, nil)
		}
	}
	return reflect.ValueOf(rvs), interpreter.ErrReturn
}
