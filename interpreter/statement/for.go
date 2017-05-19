package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// For defines a for statement.
type For struct {
	interpreter.PosImpl
	Var   string
	Value interpreter.Expr
	Stmts []interpreter.Stmt
}

func (stmt *For) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "for %s in %v {", stmt.Var, stmt.Value)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *For) Execute(env *interpreter.Env) (reflect.Value, error) {
	val, ee := stmt.Value.Invoke(env)
	if ee != nil {
		return val, ee
	}
	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		newEnv := env.NewEnv()
		defer newEnv.Destroy()

		for i := 0; i < val.Len(); i++ {
			iv := val.Index(i)
			if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
				iv = iv.Elem()
			}
			newEnv.Define(stmt.Var, iv)
			rv, err := newEnv.Run(stmt.Stmts)
			if err != nil {
				if err == interpreter.ErrBreak {
					break
				}
				if err == interpreter.ErrContinue {
					continue
				}
				if err == interpreter.ErrReturn {
					return rv, err
				}
				return rv, interpreter.NewError(stmt, err)
			}
		}
		return interpreter.NilValue, nil
	}
	return interpreter.NilValue, interpreter.NewStringError(stmt, "Invalid operation for non-array value")
}
