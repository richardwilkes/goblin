package goblin

import (
	"reflect"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/interpreter/parser"
)

// ParseAndRun parses and runs the source.
func ParseAndRun(src string) (reflect.Value, error) {
	return ParseAndRunWithEnv(src, interpreter.NewEnv())
}

// ParseAndRunWithEnv parses and runs the source in the specified environment.
func ParseAndRunWithEnv(src string, env *interpreter.Env) (reflect.Value, error) {
	stmts, err := parser.Parse(src)
	if err != nil {
		return interpreter.NilValue, err
	}
	return env.Run(stmts)
}
