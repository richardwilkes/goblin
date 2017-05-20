package goblin

import (
	"reflect"
	"testing"

	"github.com/richardwilkes/goblin/ast"
)

func TestGet(t *testing.T) {
	scope := NewScope()
	scope.Define("foo", "bar")

	v, err := scope.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != "bar" {
		t.Fatalf("Expected %v, but %v:", "bar", v.String())
	}
}

func TestDefine(t *testing.T) {
	scope := NewScope()
	scope.Define("foo", "bar")
	sub := scope.NewScope()

	v, err := sub.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != "bar" {
		t.Fatalf("Expected %v, but %v:", "bar", v.String())
	}
}

func TestDefineModify(t *testing.T) {
	scope := NewScope()
	scope.Define("foo", "bar")
	sub := scope.NewScope()
	sub.Define("foo", true)

	v, err := sub.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.Bool {
		t.Fatalf(`Can't Get bool value for "foo"`)
	}
	if !v.Bool() {
		t.Fatalf("Expected %v, but %v:", true, v.Bool())
	}

	v, err = scope.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != "bar" {
		t.Fatalf("Expected %v, but %v:", "bar", v.String())
	}
}

func TestDefineType(t *testing.T) {
	scope := NewScope()
	scope.DefineType("int", int(0))
	sub := scope.NewScope()
	sub.DefineType("str", "")
	pkg := scope.NewPackage("pkg")
	pkg.DefineType("Bool", true)

	for _, e := range []ast.Scope{scope, sub, pkg} {
		typ, err := e.Type("int")
		if err != nil {
			t.Fatalf(`Can't get Type for "int"`)
		}
		if typ.Kind() != reflect.Int {
			t.Fatalf(`Can't get int Type for "int"`)
		}

		typ, err = e.Type("str")
		if err != nil {
			t.Fatalf(`Can't get Type for "str"`)
		}
		if typ.Kind() != reflect.String {
			t.Fatalf(`Can't get string Type for "str"`)
		}

		typ, err = e.Type("pkg.Bool")
		if err != nil {
			t.Fatalf(`Can't get Type for "pkg.Bool"`)
		}
		if typ.Kind() != reflect.Bool {
			t.Fatalf(`Can't get bool Type for "pkg.Bool"`)
		}
	}
}
