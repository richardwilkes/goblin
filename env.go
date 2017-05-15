package goblin

import (
	"fmt"
	"reflect"
	"strings"
	"sync/atomic"
)

// Env holds the environment for interpreting a script.
type Env struct {
	name      string
	env       map[string]reflect.Value
	typ       map[string]reflect.Type
	parent    *Env
	interrupt *int32 // Using sync.LoadInt32 & sync.StoreInt32, so can't be bool
}

// NewEnv creates new global scope.
func NewEnv() *Env {
	var b int32
	env := &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    nil,
		interrupt: &b,
	}
	env.loadBuiltins()
	return env
}

// NewEnv creates new child scope.
func (env *Env) NewEnv() *Env {
	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    env,
		name:      env.name,
		interrupt: env.interrupt,
	}
}

// NewPackage creates a new global package.
func NewPackage(n string) *Env {
	var b int32
	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    nil,
		name:      n,
		interrupt: &b,
	}
}

// NewPackage creates a new child package.
func (env *Env) NewPackage(n string) *Env {
	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
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
	for k, v := range env.parent.env {
		if v.IsValid() && v.Interface() == env {
			delete(env.parent.env, k)
		}
	}
	env.parent = nil
	env.env = nil
}

// NewModule creates new module scope as global.
func (env *Env) NewModule(n string) *Env {
	m := &Env{
		env:    make(map[string]reflect.Value),
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
	if v, ok := env.env[sym]; ok {
		return v.Addr(), nil
	}
	if env.parent == nil {
		return NilValue, fmt.Errorf("Undefined symbol '%s'", sym)
	}
	return env.parent.Addr(sym)
}

// Type returns type which specified symbol.
func (env *Env) Type(sym string) (reflect.Type, error) {
	if v, ok := env.typ[sym]; ok {
		return v, nil
	}
	if env.parent == nil {
		return NilType, fmt.Errorf("Undefined type '%s'", sym)
	}
	return env.parent.Type(sym)
}

// Get returns value which specified symbol.
func (env *Env) Get(sym string) (reflect.Value, error) {
	if v, ok := env.env[sym]; ok {
		return v, nil
	}
	if env.parent == nil {
		return NilValue, fmt.Errorf("Undefined symbol '%s'", sym)
	}
	return env.parent.Get(sym)
}

// Set the symbol's value.
func (env *Env) Set(k string, v interface{}) error {
	if _, ok := env.env[k]; ok {
		val, ok := v.(reflect.Value)
		if !ok {
			val = reflect.ValueOf(v)
		}
		env.env[k] = val
		return nil
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

	global.typ[strings.Join(keys, ".")] = typ
}

// Define a symbol in the current scope.
func (env *Env) Define(k string, v interface{}) {
	val, ok := v.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(v)
	}
	env.env[k] = val
}

// String returns the name of current scope.
func (env *Env) String() string {
	return env.name
}

// Dump shows symbol values in the scope.
func (env *Env) Dump() {
	for k, v := range env.env {
		fmt.Printf("%v = %#v\n", k, v)
	}
}

// ParseAndRun parses and runs source in current scope.
func (env *Env) ParseAndRun(src string) (reflect.Value, error) {
	stmts, err := ParseSrc(src)
	if err != nil {
		return NilValue, err
	}
	return env.Run(stmts)
}

// Interrupt the execution of any running statements in this environment.
//
// Note that the execution is not instantly aborted: after a call to Interrupt,
// the current running statement will finish, but the next statement will not run,
// and instead will return a NilValue and an ErrInterrupt.
func (env *Env) Interrupt() {
	atomic.StoreInt32(env.interrupt, 1)
}

// RunSingleStmt executes one statement in this environment.
func (env *Env) RunSingleStmt(stmt Stmt) (reflect.Value, error) {
	if atomic.LoadInt32(env.interrupt) != 0 {
		atomic.StoreInt32(env.interrupt, 0)
		return NilValue, ErrInterrupt
	}
	return stmt.Execute(env)
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
