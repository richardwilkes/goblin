package expression

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Map defines a map expression.
type Map struct {
	ast.PosImpl
	Map map[string]ast.Expr
}

func (expr *Map) String() string {
	keys := make([]string, 0, len(expr.Map))
	for k := range expr.Map {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buffer bytes.Buffer
	buffer.WriteString("{")
	switch len(keys) {
	case 0:
	case 1:
		fmt.Fprintf(&buffer, "%s: %s", util.QuotedString(keys[0]), expr.Map[keys[0]])
	default:
		prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
		for _, k := range keys {
			fmt.Fprintf(prefixer, "\n%s: %s,", util.QuotedString(k), expr.Map[k])
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Map) Invoke(scope ast.Scope) (reflect.Value, error) {
	m := make(map[string]interface{})
	for k, e := range expr.Map {
		v, err := e.Invoke(scope)
		if err != nil {
			return v, ast.NewError(e, err)
		}
		m[k] = v.Interface()
	}
	return reflect.ValueOf(m), nil
}

// Assign a value to the expression and return it.
func (expr *Map) Assign(rv reflect.Value, scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
