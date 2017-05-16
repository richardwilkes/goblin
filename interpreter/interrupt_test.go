package interpreter_test

import (
	"sync"
	"testing"
	"time"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/interpreter/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterrupt(t *testing.T) {
	stmts, err := parser.Parse(`sleep("5s"); println("This should not be printed")`)
	require.NoError(t, err)
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go testInterrupt(t, stmts, &wg)
	}
	wg.Wait()
}

func testInterrupt(t *testing.T, stmts []interpreter.Stmt, wg *sync.WaitGroup) {
	defer wg.Done()
	env := interpreter.NewEnv()
	go func() {
		time.Sleep(time.Second)
		env.Interrupt()
	}()
	_, err := env.Run(stmts)
	assert.Equal(t, interpreter.ErrInterrupt, err)
}
