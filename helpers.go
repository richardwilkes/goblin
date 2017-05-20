package goblin

import (
	"reflect"
	"time"
)

// ParseAndRun parses and runs the script.
func ParseAndRun(script string) (reflect.Value, error) {
	return NewScope().ParseAndRun(script)
}

// ParseAndRunWithTimeout parses and runs the script and interrupts execution if the run
// time exceeds the specified timeout value. The timeout does not include the time
// required to parse the script.
func ParseAndRunWithTimeout(timeout time.Duration, script string) (reflect.Value, error) {
	return NewScope().ParseAndRunWithTimeout(timeout, script)
}
