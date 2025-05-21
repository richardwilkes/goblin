// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ast

import "reflect"

// Constants
var (
	NilValue   = reflect.ValueOf((*any)(nil))
	NilType    = reflect.TypeOf((*any)(nil))
	TrueValue  = reflect.ValueOf(true)
	FalseValue = reflect.ValueOf(false)
)
