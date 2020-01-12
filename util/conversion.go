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
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ToString converts a reflect.Value to a string.
func ToString(v reflect.Value) string {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.String {
		return v.String()
	}
	if !v.IsValid() {
		return "nil"
	}
	return fmt.Sprint(v.Interface())
}

// ToBool converts a reflect.Value to a bool.
func ToBool(v reflect.Value) bool {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0.0
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uint16:
		return v.Int() != 0
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		str := v.String()
		if len(str) < 5 {
			str = strings.ToLower(str)
			if str == "true" || str == "y" || str == "yes" {
				return true
			}
		}
		if i, err := StrToInt64(str); err == nil && i != 0 {
			return true
		}
	}
	return false
}

// ToInt64 converts a reflect.Value to an int64.
func ToInt64(v reflect.Value) int64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return int64(v.Float())
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uint16:
		return v.Int()
	case reflect.String:
		if i, err := StrToInt64(v.String()); err == nil {
			return i
		}
	case reflect.Bool:
		if v.Bool() {
			return 1
		}
	}
	return 0
}

// StrToInt64 attempts to convert a string to an int64.
func StrToInt64(str string) (int64, error) {
	if strings.HasPrefix(str, "0x") {
		return strconv.ParseInt(str[2:], 16, 64)
	}
	if i, err := strconv.ParseInt(str, 10, 64); err == nil {
		return i, nil
	}
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return int64(f), nil
	}
	return 0, err
}

// ToFloat64 converts a reflect.Value to a float64.
func ToFloat64(v reflect.Value) float64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uint16:
		return float64(v.Int())
	case reflect.String:
		str := v.String()
		if strings.HasPrefix(str, "0x") {
			if i, err := strconv.ParseInt(str[2:], 16, 64); err == nil {
				return float64(i)
			}
		} else {
			if i, err := strconv.ParseInt(str, 10, 64); err == nil {
				return float64(i)
			}
			if f, err := strconv.ParseFloat(str, 64); err == nil {
				return f
			}
		}
	case reflect.Bool:
		if v.Bool() {
			return 1.0
		}
	}
	return 0.0
}

// ToSlice converts a generic slice to a typed slice.
func ToSlice(from []interface{}, ptr interface{}) {
	obj := reflect.Indirect(reflect.ValueOf(ptr))
	slice := reflect.MakeSlice(reflect.TypeOf(ptr).Elem(), len(from), len(from))
	for i, v := range from {
		slice.Index(i).Set(reflect.ValueOf(v))
	}
	obj.Set(slice)
}
