// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package parser_test

import (
	"testing"

	"github.com/richardwilkes/goblin/parser"
	"github.com/richardwilkes/toolbox/check"
)

func TestParsingError(t *testing.T) {
	script := `x3.4`
	_, err := parser.Parse(script)
	check.Error(t, err, script)
	check.Contains(t, err.Error(), "1:4", script)

	script = `
for {
	i = 1
	j = 2 k = 3
}`
	_, err = parser.Parse(script)
	check.Error(t, err, script)
	check.Contains(t, err.Error(), "4:8", script)
}
