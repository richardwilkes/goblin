package goblin

import (
	"fmt"
	"reflect"
)

// Func defines internal functions.
type Func func(args ...reflect.Value) (reflect.Value, error)

func (f Func) String() string {
	return fmt.Sprintf("[Function: %p]", f)
}
