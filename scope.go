package goblin

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/parser"
)

type scope struct {
	name      string
	env       map[string]reflect.Value // Only created when needed to reduce memory thrashing
	typ       map[string]reflect.Type  // Only created when needed to reduce memory thrashing
	parent    *scope
	interrupt *chan bool
}

// NewScope creates new global scope.
func NewScope() ast.Scope {
	interrupt := make(chan bool, 2)
	s := &scope{interrupt: &interrupt}
	s.loadBuiltins()
	return s
}

func (s *scope) NewScope() ast.Scope {
	return &scope{
		parent:    s,
		name:      s.name,
		interrupt: s.interrupt,
	}
}

func (s *scope) NewPackage(n string) ast.Scope {
	return &scope{
		parent:    s,
		name:      n,
		interrupt: s.interrupt,
	}
}

func (s *scope) NewModule(n string) ast.Scope {
	m := &scope{
		parent: s,
		name:   n,
	}
	s.Define(n, m)
	return m
}

func (s *scope) Run(stmts []ast.Stmt) (reflect.Value, error) {
	s.drainInterrupts()
	return s.run(stmts)
}

func (s *scope) RunWithTimeout(timeout time.Duration, stmts []ast.Stmt) (reflect.Value, error) {
	if timeout < 1 {
		return ast.NilValue, ast.ErrInterrupt
	}
	s.drainInterrupts()
	timer := time.AfterFunc(timeout, s.Interrupt)
	v, err := s.run(stmts)
	timer.Stop()
	return v, err
}

func (s *scope) ParseAndRun(script string) (reflect.Value, error) {
	stmts, err := parser.Parse(script)
	if err != nil {
		return ast.NilValue, err
	}
	return s.Run(stmts)
}

func (s *scope) ParseAndRunWithTimeout(timeout time.Duration, script string) (reflect.Value, error) {
	stmts, err := parser.Parse(script)
	if err != nil {
		return ast.NilValue, err
	}
	return s.RunWithTimeout(timeout, stmts)
}

func (s *scope) drainInterrupts() {
	if s.parent == nil {
	out:
		for {
			select {
			case <-*(s.interrupt):
			default:
				break out
			}
		}
	}
}

func (s *scope) run(stmts []ast.Stmt) (reflect.Value, error) {
	rv := ast.NilValue
	var err error
	for _, stmt := range stmts {
		select {
		case <-*(s.interrupt):
			return ast.NilValue, ast.ErrInterrupt
		default:
			if rv, err = stmt.Execute(s); err != nil {
				return rv, err
			}
		}
	}
	return rv, nil
}

func (s *scope) Interrupt() {
	*(s.interrupt) <- true
}

func (s *scope) Destroy() {
	if s.parent != nil {
		if s.parent.env != nil {
			for k, v := range s.parent.env {
				if v.IsValid() && v.Interface() == s {
					delete(s.parent.env, k)
				}
			}
		}
		s.parent = nil
		s.env = nil
	}
}

func (s *scope) Type(sym string) (reflect.Type, error) {
	if s.typ != nil {
		if v, ok := s.typ[sym]; ok {
			return v, nil
		}
	}
	if s.parent == nil {
		return ast.NilType, fmt.Errorf("Undefined type '%s'", sym)
	}
	return s.parent.Type(sym)
}

func (s *scope) Get(sym string) (reflect.Value, error) {
	if s.env != nil {
		if v, ok := s.env[sym]; ok {
			return v, nil
		}
	}
	if s.parent == nil {
		return ast.NilValue, fmt.Errorf("Undefined symbol '%s'", sym)
	}
	return s.parent.Get(sym)
}

func (s *scope) Set(k string, v interface{}) error {
	if s.env != nil {
		if _, ok := s.env[k]; ok {
			val, ok := v.(reflect.Value)
			if !ok {
				val = reflect.ValueOf(v)
			}
			s.env[k] = val
			return nil
		}
	}
	if s.parent == nil {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	return s.parent.Set(k, v)
}

func (s *scope) DefineGlobal(k string, v interface{}) {
	if s.parent == nil {
		s.Define(k, v)
	} else {
		s.parent.DefineGlobal(k, v)
	}
}

func (s *scope) DefineType(k string, t interface{}) {
	global := s
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

func (s *scope) Define(k string, v interface{}) {
	val, ok := v.(reflect.Value)
	if !ok {
		val = reflect.ValueOf(v)
	}
	if s.env == nil {
		s.env = make(map[string]reflect.Value)
	}
	s.env[k] = val
}

func (s *scope) String() string {
	return fmt.Sprintf("[scope: %s]", s.name)
}
