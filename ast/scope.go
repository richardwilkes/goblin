// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ast

import (
	"reflect"
	"time"
)

// Scope provides scoping of symbols when executing a script.
type Scope interface {
	// NewScope creates new child scope.
	NewScope() Scope
	// NewPackage creates a new child package scope.
	NewPackage(name string) Scope
	// NewModule creates new child module scope.
	NewModule(name string) Scope
	// Run executes statements in this scope.
	Run(stmts []Stmt) (reflect.Value, error)
	// RunWithTimeout executes statements in this scope and interrupts execution if
	// the run time exceeds the specified timeout value.
	RunWithTimeout(timeout time.Duration, stmts []Stmt) (reflect.Value, error)
	// ParseAndRun parses and runs the script in this scope.
	ParseAndRun(script string) (reflect.Value, error)
	// ParseAndRunWithTimeout parses and runs the script in this scope and interrupts
	// execution if the run time exceeds the specified timeout value. The timeout does
	// not include the time required to parse the script.
	ParseAndRunWithTimeout(timeout time.Duration, script string) (reflect.Value, error)
	// Interrupt the execution of any running statements in this scope or its children.
	//
	// Note that the execution is not instantly aborted. After a call to Interrupt(),
	// the current running statement will finish, but the next statement will not run,
	// and instead will return a NilValue and an ErrInterrupt.
	Interrupt()
	// Destroy deletes this scope from its parent, if any.
	Destroy()
	// Type returns the type of the specified symbol.
	Type(sym string) (reflect.Type, error)
	// Get returns value of the specified symbol.
	Get(sym string) (reflect.Value, error)
	// Set the symbol's value.
	Set(k string, v interface{}) error
	// DefineGlobal defines a symbol in the global scope.
	DefineGlobal(k string, v interface{})
	// DefineType defines a type.
	DefineType(k string, t interface{})
	// Define a symbol in this scope.
	Define(k string, v interface{})
}
