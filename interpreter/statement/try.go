package statement

import (
	"fmt"
	"reflect"

	"bytes"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// Try defines the try/catch/finally statement.
type Try struct {
	interpreter.PosImpl
	Try     []interpreter.Stmt
	Var     string
	Catch   []interpreter.Stmt
	Finally []interpreter.Stmt
}

func (stmt *Try) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("try {")
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Try {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n} catch ")
	if stmt.Var != "" {
		buffer.WriteString(stmt.Var)
		buffer.WriteString(" ")
	}
	buffer.WriteString("{")
	for _, one := range stmt.Catch {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	if len(stmt.Finally) > 0 {
		buffer.WriteString(" finally {")
		for _, one := range stmt.Finally {
			fmt.Fprintf(prefixer, "\n%v", one)
		}
		buffer.WriteString("\n}")
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Try) Execute(env *interpreter.Env) (reflect.Value, error) {
	newEnv := env.NewEnv()
	defer newEnv.Destroy()
	_, err := newEnv.Run(stmt.Try)
	if err != nil {
		// Catch
		cenv := env.NewEnv()
		defer cenv.Destroy()
		if stmt.Var != "" {
			cenv.Define(stmt.Var, reflect.ValueOf(err))
		}
		_, e1 := cenv.Run(stmt.Catch)
		if e1 != nil {
			err = interpreter.NewError(stmt.Catch[0], e1)
		} else {
			err = nil
		}
	}
	if len(stmt.Finally) > 0 {
		// Finally
		fenv := env.NewEnv()
		defer fenv.Destroy()
		_, e2 := fenv.Run(stmt.Finally)
		if e2 != nil {
			err = interpreter.NewError(stmt.Finally[0], e2)
		}
	}
	return interpreter.NilValue, interpreter.NewError(stmt, err)
}
