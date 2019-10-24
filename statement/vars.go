package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Variables defines a statement which defines multiple variables.
type Variables struct {
	ast.PosImpl
	Left     []ast.Expr
	Operator string
	Right    []ast.Expr
}

func (stmt *Variables) String() string {
	var buffer bytes.Buffer
	for i, one := range stmt.Left {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	buffer.WriteString(" ")
	buffer.WriteString(stmt.Operator)
	buffer.WriteString(" ")
	for i, one := range stmt.Right {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Variables) Execute(scope ast.Scope) (reflect.Value, error) {
	rv := ast.NilValue
	var err error
	vs := make([]interface{}, 0, len(stmt.Right))
	for _, right := range stmt.Right {
		rv, err = right.Invoke(scope)
		if err != nil {
			return rv, ast.NewError(right, err)
		}
		switch {
		case rv == ast.NilValue:
			vs = append(vs, nil)
		case rv.IsValid() && rv.CanInterface():
			vs = append(vs, rv.Interface())
		default:
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
		_, err = left.Assign(v, scope)
		if err != nil {
			return rvs, ast.NewError(left, err)
		}
	}
	if rvs.Len() == 1 {
		return rvs.Index(0), nil
	}
	return rvs, nil
}
