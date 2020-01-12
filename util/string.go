// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package util

import "bytes"

// QuotedString returns a quoted string suitable for a script statement.
func QuotedString(str string) string {
	var buffer bytes.Buffer
	buffer.WriteString(`"`)
	for _, ch := range str {
		switch ch {
		case '"':
			buffer.WriteString(`\"`)
		case '\n':
			buffer.WriteString(`\n`)
		case '\t':
			buffer.WriteString(`\t`)
		default:
			buffer.WriteRune(ch)
		}
	}
	buffer.WriteString(`"`)
	return buffer.String()
}
