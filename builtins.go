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
	"fmt"
	"reflect"
	"sort"
	"time"
	"unicode/utf8"

	"github.com/richardwilkes/goblin/util"
)

func (s *scope) loadBuiltins() {
	s.Define("len", func(v any) int64 {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		k := rv.Kind()
		if k == reflect.String || k == reflect.Array || k == reflect.Slice || k == reflect.Map {
			return int64(rv.Len())
		}
		return 0
	})

	s.Define("keys", func(v any) []any {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Map {
			k := rv.MapKeys()
			keys := make([]any, len(k))
			for i := range k {
				keys[i] = k[i].Interface()
			}
			return keys
		}
		return []any{}
	})

	s.Define("range", func(args ...int64) []int64 {
		count := len(args)
		if count == 0 {
			return []int64{}
		}
		var minimum, maximum int64
		if count == 1 {
			maximum = args[0] - 1
			if maximum < 0 {
				return []int64{}
			}
		} else {
			minimum = args[0]
			maximum = args[1]
		}
		if maximum < minimum {
			maximum = minimum
		}
		arr := make([]int64, 1+maximum-minimum)
		for i := minimum; i <= maximum; i++ {
			arr[i-minimum] = i
		}
		return arr
	})

	s.Define("sort", sort.Slice)

	s.Define("sleep", func(spec string) {
		if d, err := time.ParseDuration(spec); err == nil && d > 0 {
			timer := time.NewTimer(d)
			select {
			case <-timer.C:
			case <-(*s.interrupt):
				timer.Stop()
				s.Interrupt() // Re-interrupt, since we just ate the previous one
			}
		}
	})

	s.Define("toString", func(v any) string {
		if str, ok := v.(string); ok {
			return str
		}
		if b, ok := v.([]byte); ok {
			return string(b)
		}
		return fmt.Sprint(v)
	})

	s.Define("toInt", func(v any) int64 {
		return util.ToInt64(reflect.ValueOf(v))
	})

	s.Define("toFloat", func(v any) float64 {
		return util.ToFloat64(reflect.ValueOf(v))
	})

	s.Define("toBool", func(v any) bool {
		return util.ToBool(reflect.ValueOf(v))
	})

	s.Define("toChar", func(s rune) string {
		return string(s)
	})

	s.Define("toRune", func(s string) rune {
		if s == "" {
			return 0
		}
		r, _ := utf8.DecodeRuneInString(s)
		return r
	})

	s.Define("toByteSlice", func(s string) []byte {
		return []byte(s)
	})

	s.Define("toRuneSlice", func(s string) []rune {
		return []rune(s)
	})

	s.Define("toBoolSlice", func(v []any) []bool {
		var result []bool
		util.ToSlice(v, &result)
		return result
	})

	s.Define("toIntSlice", func(v []any) []int64 {
		var result []int64
		util.ToSlice(v, &result)
		return result
	})

	s.Define("toFloatSlice", func(v []any) []float64 {
		var result []float64
		util.ToSlice(v, &result)
		return result
	})

	s.Define("toStringSlice", func(v []any) []string {
		var result []string
		util.ToSlice(v, &result)
		return result
	})

	s.Define("typeOf", func(v any) string {
		t := reflect.TypeOf(v)
		if t == nil {
			return "<nil>"
		}
		return reflect.TypeOf(v).String()
	})

	s.Define("defined", func(str string) bool {
		_, err := s.Get(str)
		return err == nil
	})

	s.Define("print", fmt.Print)
	s.Define("println", fmt.Println)
	s.Define("printf", fmt.Printf)

	s.DefineType("int", int64(0))
	s.DefineType("float", 0.0)
	s.DefineType("bool", true)
	s.DefineType("string", "")
}
