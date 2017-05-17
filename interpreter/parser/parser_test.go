package parser_test

import (
	"testing"

	"github.com/richardwilkes/goblin/interpreter/parser"
	"github.com/stretchr/testify/assert"
)

func TestParsingError(t *testing.T) {
	script := `x3.4`
	_, err := parser.Parse(script)
	assert.Error(t, err, script)
	assert.Contains(t, err.Error(), "1:4", script)

	script = `
for {
	i = 1
	j = 2 k = 3
}`
	_, err = parser.Parse(script)
	assert.Error(t, err, script)
	assert.Contains(t, err.Error(), "4:8", script)
}
