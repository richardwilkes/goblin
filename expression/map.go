// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package expression

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/richardwilkes/goblin/ast"
	"github.com/richardwilkes/goblin/util"
)

// Map defines a map expression.
type Map struct {
	Map map[string]ast.Expr
	ast.PosImpl
}

func (expr *Map) String() string {
	keys := make([]string, 0, len(expr.Map))
	for k := range expr.Map {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var buffer bytes.Buffer
	buffer.WriteString("{")
	switch len(keys) {
	case 0:
	case 1:
		fmt.Fprintf(&buffer, "%s: %s", util.QuotedString(keys[0]), expr.Map[keys[0]])
	default:
		prefixer := &util.Prefixer{Prefix: "    ", Writer: &buffer}
		for _, k := range keys {
			fmt.Fprintf(prefixer, "\n%s: %s,", util.QuotedString(k), expr.Map[k])
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Invoke the expression and return a result.
func (expr *Map) Invoke(scope ast.Scope) (reflect.Value, error) {
	m := make(map[string]any)
	for k, e := range expr.Map {
		v, err := e.Invoke(scope)
		if err != nil {
			return v, ast.NewError(e, err)
		}
		m[k] = v.Interface()
	}
	return reflect.ValueOf(m), nil
}

// Assign a value to the expression and return it.
func (expr *Map) Assign(_ reflect.Value, _ ast.Scope) (reflect.Value, error) {
	return ast.NilValue, ast.NewInvalidOperationError(expr)
}
