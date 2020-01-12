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

// Case defines a case statement.
type Case struct {
	ast.PosImpl
	Expr  ast.Expr
	Stmts []ast.Stmt
}

func (stmt *Case) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "case %v:", stmt.Expr)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, s := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", s)
	}
	return buffer.String()
}

// Execute the statement.
func (stmt *Case) Execute(scope ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewCannotExecuteError(stmt)
}
