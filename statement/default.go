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
	"bytes"
	"fmt"
	"reflect"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Default defines the default case statement.
type Default struct {
	Stmts []ast.Stmt
	ast.PosImpl
}

func (stmt *Default) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("default:")
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Default) Execute(_ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewCannotExecuteError(stmt)
}
