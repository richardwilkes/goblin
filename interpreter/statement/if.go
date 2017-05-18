package statement

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// If defines an if/else statement.
type If struct {
	interpreter.PosImpl
	If     interpreter.Expr
	Then   []interpreter.Stmt
	ElseIf []interpreter.Stmt
	Else   []interpreter.Stmt
}

func (stmt *If) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "if %v {\n", stmt.If)
	for _, one := range stmt.Then {
		fmt.Fprintf(&buffer, "    %v\n", one)
	}
	buffer.WriteString("}")
	if len(stmt.ElseIf) > 0 {
		for _, one := range stmt.ElseIf {
			fmt.Fprintf(&buffer, " else %v", one)
		}
	}
	if len(stmt.Else) > 0 {
		buffer.WriteString(" else {\n")
		for _, one := range stmt.Else {
			fmt.Fprintf(&buffer, "    %v\n", one)
		}
		buffer.WriteString("}")
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *If) Execute(env *interpreter.Env) (reflect.Value, error) {
	// If
	rv, err := stmt.If.Invoke(env)
	if err != nil {
		return rv, interpreter.NewError(stmt, err)
	}
	if util.ToBool(rv) {
		// Then
		newEnv := env.NewEnv()
		defer newEnv.Destroy()
		rv, err = newEnv.Run(stmt.Then)
		if err != nil {
			return rv, interpreter.NewError(stmt, err)
		}
		return rv, nil
	}
	done := false
	if len(stmt.ElseIf) > 0 {
		for _, stmt := range stmt.ElseIf {
			stmtIf, ok := stmt.(*If)
			if !ok {
				return interpreter.NilValue, interpreter.NewError(stmt, interpreter.ErrBadSyntax)
			}
			// Else If
			rv, err = stmtIf.If.Invoke(env)
			if err != nil {
				return rv, interpreter.NewError(stmt, err)
			}
			if !util.ToBool(rv) {
				continue
			}
			// Else If Then
			done = true
			rv, err = env.Run(stmtIf.Then)
			if err != nil {
				return rv, interpreter.NewError(stmt, err)
			}
			break
		}
	}
	if !done && len(stmt.Else) > 0 {
		// Else
		newEnv := env.NewEnv()
		defer newEnv.Destroy()
		rv, err = newEnv.Run(stmt.Else)
		if err != nil {
			return rv, interpreter.NewError(stmt, err)
		}
	}
	return rv, nil
}
