// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package util

import (
	"io"
	"unicode/utf8"
)

// Prefixer provides an io.Writer that will add a prefix to each line output.
type Prefixer struct {
	Prefix string
	Writer io.Writer
}

func (prefixer *Prefixer) Write(p []byte) (n int, err error) {
	if prefixer.Prefix == "" {
		return prefixer.Writer.Write(p)
	}
	prefix := []byte(prefixer.Prefix)
	var n1 int
	i := 0
	j := 0
	for j < len(p) {
		r, size := utf8.DecodeRune(p[j:])
		if r == '\n' {
			j++
			n1, err = prefixer.Writer.Write(p[i:j])
			n += n1
			if err != nil {
				return
			}
			n1, err = prefixer.Writer.Write(prefix)
			n += n1
			if err != nil {
				return
			}
			i = j
		} else {
			j += size
		}
	}
	if i != j {
		n1, err = prefixer.Writer.Write(p[i:j])
		n += n1
	}
	return
}
