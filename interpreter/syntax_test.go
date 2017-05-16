package interpreter_test

import (
	"testing"

	"github.com/richardwilkes/goblin/interpreter/parser"
	"github.com/stretchr/testify/assert"
)

func TestComments(t *testing.T) {
	checkSyntax(t, "# single line comment")
	checkSyntax(t, "x = 1 # single line comment")
	checkSyntax(t, "x = 1# single line comment")
	checkSyntax(t, "// single line comment")
	checkSyntax(t, "x = 1 // single line comment")
	checkSyntax(t, "x = 1// single line comment")
	checkSyntax(t, `/*
 * multi-line comment
 */`)
	checkSyntax(t, `x = 1/*
 * multi-line comment
 */
y = 2`)
	checkSyntax(t, `for i = 0; i < 10 /* comment */; i++ {}`)
}

func checkSyntax(t *testing.T, script string) {
	_, err := parser.Parse(script)
	assert.NoError(t, err, script)
}
