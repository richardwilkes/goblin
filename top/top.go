package top

import (
	"reflect"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/parser"
)

// ParseAndRun parses and runs source in current scope.
func ParseAndRun(src string, env *goblin.Env) (reflect.Value, error) {
	stmts, err := parser.ParseSrc(src)
	if err != nil {
		return goblin.NilValue, err
	}
	return env.Run(stmts)
}
