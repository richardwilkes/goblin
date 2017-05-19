package statement

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Loop defines a loop statement.
type Loop struct {
	interpreter.PosImpl
	Expr  interpreter.Expr
	Stmts []interpreter.Stmt
}

func (stmt *Loop) String() string {
	var buffer bytes.Buffer
	if stmt.Expr != nil {
		fmt.Fprintf(&buffer, "for %v {", stmt.Expr)
	} else {
		buffer.WriteString("for {")
	}
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *Loop) Execute(env *interpreter.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	for {
		if stmt.Expr != nil {
			ev, ee := stmt.Expr.Invoke(newEnv)
			if ee != nil {
				return ev, ee
			}
			if !util.ToBool(ev) {
				break
			}
		}

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
