// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package goblin_test

import (
	"reflect"
	"testing"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/ast"
)

const (
	foo = "foo"
	bar = "bar"
)

func TestGet(t *testing.T) {
	scope := goblin.NewScope()
	scope.Define(foo, bar)

	v, err := scope.Get(foo)
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != bar {
		t.Fatalf("Expected %v, but %v:", bar, v.String())
	}
}

func TestDefine(t *testing.T) {
	scope := goblin.NewScope()
	scope.Define(foo, bar)
	sub := scope.NewScope()

	v, err := sub.Get(foo)
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != bar {
		t.Fatalf("Expected %v, but %v:", bar, v.String())
	}
}

func TestDefineModify(t *testing.T) {
	scope := goblin.NewScope()
	scope.Define(foo, bar)
	sub := scope.NewScope()
	sub.Define(foo, true)

	v, err := sub.Get(foo)
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.Bool {
		t.Fatalf(`Can't Get bool value for "foo"`)
	}
	if !v.Bool() {
		t.Fatalf("Expected %v, but %v:", true, v.Bool())
	}

	v, err = scope.Get(foo)
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != bar {
		t.Fatalf("Expected %v, but %v:", bar, v.String())
	}
}

func TestDefineType(t *testing.T) {
	scope := goblin.NewScope()
	scope.DefineType("int", 0)
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
