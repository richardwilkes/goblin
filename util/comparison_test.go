// Copyright ©2017-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package util_test

import (
	"reflect"
	"testing"

	"github.com/richardwilkes/goblin/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsNil(t *testing.T) {
	input := []interface{}{nil, "1", 1, 1.0, &[]string{}, &[]string{"hello"}}
	result := []bool{true, false, false, false, false, false}
	require.Equal(t, len(input), len(result))
	for i := range input {
		assert.Equal(t, result[i], util.IsNil(reflect.ValueOf(input[i])), "%v", input[i])
	}
}

func TestIsNumber(t *testing.T) {
	input := []interface{}{nil, "1", 2, 3.0, &[]string{}, &[]string{"hello"}, "goodbye"}
	result := []bool{false, false, true, true, false, false, false}
	require.Equal(t, len(input), len(result))
	for i := range input {
		assert.Equal(t, result[i], util.IsNumber(reflect.ValueOf(input[i])), "%v", input[i])
	}
}

func TestEqual(t *testing.T) {
	left := []interface{}{nil, "1", 2, 3.0, &[]string{}, &[]string{"hello"}, "goodbye"}
	right := []interface{}{nil, "1", 2, 2.0, 3, 3.0, &[]string{}, &[]string{"hello"}, "goodbye"}
	result := [][]bool{
		{true, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false},
		{false, false, true, true, false, false, false, false, false},
		{false, false, false, false, true, true, false, false, false},
		{false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, true, false},
		{false, false, false, false, false, false, false, false, true},
	}
	require.Equal(t, len(left), len(result))
	for _, one := range result {
		require.Equal(t, len(right), len(one))
	}
	for i := range left {
		for j := range right {
			assert.Equal(t, result[i][j], util.Equal(reflect.ValueOf(left[i]), reflect.ValueOf(right[j])), "%v == %v", left[i], right[i])
		}
	}
}
