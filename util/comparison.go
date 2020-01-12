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
	"reflect"
)

// IsNil returns true if the reflect.Value is equivalent to nil.
func IsNil(v reflect.Value) bool {
	if !v.IsValid() || v.Kind().String() == "unsafe.Pointer" {
		return true
	}
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.IsNil() {
		return true
	}
	return false
}

// IsNumber returns true if the reflect.Value is a number value.
func IsNumber(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// Equal returns true if the two reflect.Values are equal.
func Equal(left, right reflect.Value) bool {
	leftIsNil, rightIsNil := IsNil(left), IsNil(right)
	if leftIsNil && rightIsNil {
		return true
	}
	if (!leftIsNil && rightIsNil) || (leftIsNil && !rightIsNil) {
		return false
	}
	if left.Kind() == reflect.Interface || left.Kind() == reflect.Ptr {
		left = left.Elem()
	}
	if right.Kind() == reflect.Interface || right.Kind() == reflect.Ptr {
		right = right.Elem()
	}
	if !left.IsValid() || !right.IsValid() {
		return true
	}
	leftNum := IsNumber(left)
	rightNum := IsNumber(right)
	if leftNum && rightNum {
		if right.Type().ConvertibleTo(left.Type()) {
			right = right.Convert(left.Type())
		}
	}
	if leftNum && !rightNum && right.Kind() == reflect.String {
		if rv, err := StrToInt64(right.String()); err == nil {
			right = reflect.ValueOf(rv)
			if right.Type().ConvertibleTo(left.Type()) {
				right = right.Convert(left.Type())
			}
		}
	}
	if rightNum && !leftNum && left.Kind() == reflect.String {
		if lv, err := StrToInt64(left.String()); err == nil {
			left = reflect.ValueOf(lv)
			if left.Type().ConvertibleTo(right.Type()) {
				left = left.Convert(right.Type())
			}
		}
	}
	if left.CanInterface() && right.CanInterface() {
		return reflect.DeepEqual(left.Interface(), right.Interface())
	}
	return reflect.DeepEqual(left, right)
}
