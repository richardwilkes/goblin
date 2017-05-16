package interpreter_test

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/interpreter/parser"
)

func testInterrupt(t *testing.T, stmts []interpreter.Stmt, wg *sync.WaitGroup) {
	defer wg.Done()

	env := interpreter.NewEnv()
	env.Define("sleep", func(spec string) {
		if d, err := time.ParseDuration(spec); err != nil {
			panic(err)
		} else {
			time.Sleep(d)
		}
	})

	// Interrupts after 1 second.
	go func() {
		time.Sleep(time.Second)
		env.Interrupt()
	}()

	_, err := env.Run(stmts)
	if err != interpreter.ErrInterrupt {
		t.Fail()
	}
}

func TestInterruptRaces(t *testing.T) {
	script := `sleep("2s")
# Should interrupt here.
# The next line will not be executed.
println("This should not be printed")`
	stmts, err := parser.Parse(script)
	if err != nil {
		log.Fatal()
	}

	var wg sync.WaitGroup
	// Run example several times
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go testInterrupt(t, stmts, &wg)
	}
	wg.Wait()
}