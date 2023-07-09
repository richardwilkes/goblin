// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package statement

import (
	"reflect"

	"github.com/richardwilkes/goblin/ast"
)

// Continue defines the continue statement.
type Continue struct {
	ast.PosImpl
}

func (stmt *Continue) String() string {
	return "continue"
}

// Execute the statement.
func (stmt *Continue) Execute(_ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.ErrContinue
}
