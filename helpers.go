// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package goblin

import (
	"reflect"
	"time"
)

// ParseAndRun parses and runs the script.
func ParseAndRun(script string) (reflect.Value, error) {
	return NewScope().ParseAndRun(script)
}

// ParseAndRunWithTimeout parses and runs the script and interrupts execution if the run
// time exceeds the specified timeout value. The timeout does not include the time
// required to parse the script.
func ParseAndRunWithTimeout(timeout time.Duration, script string) (reflect.Value, error) {
	return NewScope().ParseAndRunWithTimeout(timeout, script)
}
