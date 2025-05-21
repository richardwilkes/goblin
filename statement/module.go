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

// Module defines a module statement.
type Module struct {
	Name  string
	Stmts []ast.Stmt
	ast.PosImpl
}

func (stmt *Module) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "module %s {", stmt.Name)
	prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
	for _, one := range stmt.Stmts {
		fmt.Fprintf(prefixer, "\n%v", one)
	}
	buffer.WriteString("\n}")
	return buffer.String()
}

// Execute the statement.
func (stmt *Module) Execute(scope ast.Scope) (reflect.Value, error) {
	newScope := scope.NewModule(stmt.Name)
	rv, err := newScope.Run(stmt.Stmts)
	if err != nil {
		return rv, ast.NewError(stmt, err)
	}
	scope.DefineGlobal(stmt.Name, reflect.ValueOf(newScope))
	return rv, nil
}
