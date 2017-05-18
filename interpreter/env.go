package interpreter

import (
	"fmt"
	"reflect"
	"strings"
)

// Env holds the environment for interpreting a script.
type Env struct {
	name      string
	env       map[string]reflect.Value // Only created when needed to reduce memory thrashing
	typ       map[string]reflect.Type  // Only created when needed to reduce memory thrashing
	parent    *Env
	interrupt *chan bool
}

// NewEnv creates new global scope.
func NewEnv() *Env {
	interrupt := make(chan bool, 1)
	env := &Env{interrupt: &interrupt}
	env.loadBuiltins()
	return env
}

// NewEnv creates new child scope.
func (env *Env) NewEnv() *Env {
	return &Env{
		parent:    env,
		name:      env.name,
		interrupt: env.interrupt,
	}
}

// NewPackage creates a new global package.
func NewPackage(n string) *Env {
	interrupt := make(chan bool, 1)
	return &Env{
		name:      n,
		interrupt: &interrupt,
	}
}

// NewPackage creates a new child package.
func (env *Env) NewPackage(n string) *Env {
	return &Env{
		parent:    env,
		name:      n,
		interrupt: env.interrupt,
	}
}

// Destroy deletes current scope.
func (env *Env) Destroy() {
	if env.parent == nil {
		return
	}
	if env.parent.env != nil {
		for k, v := range env.parent.env {
			if v.IsValid() && v.Interface() == env {
				delete(env.parent.env, k)
			}
		}
	}
	env.parent = nil
	env.env = nil
}

// NewModule creates new module scope as global.
func (env *Env) NewModule(n string) *Env {
	m := &Env{
		parent: env,
		name:   n,
	}
	env.Define(n, m)
	return m
}

// SetName sets a name of the scope. This means that the scope is module.
func (env *Env) SetName(n string) {
	env.name = n
}

// GetName returns module name.
func (env *Env) GetName() string {
	return env.name
}

// Addr returns pointer value which specified the symbol.
func (env *Env) Addr(sym string) (reflect.Value, error) {
	if env.env != nil {
		if v, ok := env.env[sym]; ok {
			return v.Addr(), nil
		}
	}
	if env.parent == nil {
		return NilValue, fmt.Errorf("Undefined symbol '%s'", sym)
	}
	return env.parent.Addr(sym)
}

// Type returns type which specified symbol.
func (env *Env) Type(sym string) (reflect.Type, error) {
	if env.typ != nil {
		if v, ok := env.typ[sym]; ok {
			return v, nil
		}
	}
	if env.parent == nil {
		return NilType, fmt.Errorf("Undefined type '%s'", sym)
	}
	return env.parent.Type(sym)
}

// Get returns value which specified symbol.
func (env *Env) Get(sym string) (reflect.Value, error) {
	if env.env != nil {
		if v, ok := env.env[sym]; ok {
			return v, nil
		}
	}
	if env.parent == nil {
		return NilValue, fmt.Errorf("Undefined symbol '%s'", sym)
	}
	return env.parent.Get(sym)
}

// Set the symbol's value.
func (env *Env) Set(k string, v interface{}) error {
	if env.env != nil {
		if _, ok := env.env[k]; ok {
			val, ok := v.(reflect.Value)
			if !ok {
				val = reflect.ValueOf(v)
			}
			env.env[k] = val
			return nil
		}
	}
	if env.parent == nil {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	return env.parent.Set(k, v)
}

// DefineGlobal defines a symbol in global scope.
func (env *Env) DefineGlobal(k string, v interface{}) {
	if env.parent == nil {
		env.Define(k, v)
	} else {
		env.parent.DefineGlobal(k, v)
	}
}

// DefineType defines a type.
func (env *Env) DefineType(k string, t interface{}) {
	global := env
	keys := []string{k}

	for global.parent != nil {
		if global.name != "" {
			keys = append(keys, global.name)
		}
		global = global.parent
	}

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	typ, ok := t.(reflect.Type)
	if !ok {
		typ = reflect.TypeOf(t)
	}
	if global.typ == nil {
		global.typ = make(map[string]reflect.Type)
	}
	global.typ[strings.Join(keys, ".")] = typ
}

// Define a symbol in the current scope.
func (env *Env) Define(k string, v interface{}) {
	val, ok := v.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(v)
	}
	if env.env == nil {
		env.env = make(map[string]reflect.Value)
	}
	env.env[k] = val
}

// String returns the name of current scope.
func (env *Env) String() string {
	return env.name
}

// Dump shows symbol values in the scope.
func (env *Env) Dump() {
	if env.env != nil {
		for k, v := range env.env {
			fmt.Printf("%v = %#v\n", k, v)
		}
	}
}

// Interrupt the execution of any running statements in this environment.
//
// Note that the execution is not instantly aborted: after a call to Interrupt,
// the current running statement will finish, but the next statement will not run,
// and instead will return a NilValue and an ErrInterrupt.
func (env *Env) Interrupt() {
	*(env.interrupt) <- true
}

// RunSingleStmt executes one statement in this environment.
func (env *Env) RunSingleStmt(stmt Stmt) (reflect.Value, error) {
	select {
	case <-*(env.interrupt):
		return NilValue, ErrInterrupt
	default:
		return stmt.Execute(env)
	}
}

// Run executes statements in this environment.
func (env *Env) Run(stmts []Stmt) (reflect.Value, error) {
	rv := NilValue
	var err error
	for _, stmt := range stmts {
		if rv, err = env.RunSingleStmt(stmt); err != nil {
			return rv, err
		}
	}
	return rv, nil
}
