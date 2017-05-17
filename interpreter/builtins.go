package interpreter

import (
	"fmt"
	"reflect"
	"time"

	"github.com/richardwilkes/goblin/util"
)

func (env *Env) loadBuiltins() {
	env.Define("len", func(v interface{}) int64 {
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

	env.Define("keys", func(v interface{}) []interface{} {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Map {
			k := rv.MapKeys()
			keys := make([]interface{}, len(k))
			for i := range k {
				keys[i] = k[i].Interface()
			}
			return keys
		}
		return []interface{}{}
	})

	env.Define("range", func(args ...int64) []int64 {
		count := len(args)
		if count == 0 {
			return []int64{}
		}
		var min, max int64
		if count == 1 {
			max = args[0] - 1
			if max < 0 {
				return []int64{}
			}
		} else {
			min = args[0]
			max = args[1]
		}
		if max < min {
			max = min
		}
		arr := make([]int64, 1+max-min)
		for i := min; i <= max; i++ {
			arr[i-min] = i
		}
		return arr
	})

	env.Define("sleep", func(spec string) {
		if d, err := time.ParseDuration(spec); err == nil && d > 0 {
			timer := time.NewTimer(d)
			select {
			case <-timer.C:
			case <-(*env.interrupt):
				timer.Stop()
				env.Interrupt() // Re-interrupt, since we just ate the previous one
			}
		}
	})

	env.Define("toString", func(v interface{}) string {
		if s, ok := v.(string); ok {
			return s
		}
		if b, ok := v.([]byte); ok {
			return string(b)
		}
		return fmt.Sprint(v)
	})

	env.Define("toInt", func(v interface{}) int64 {
		return util.ToInt64(reflect.ValueOf(v))
	})

	env.Define("toFloat", func(v interface{}) float64 {
		return util.ToFloat64(reflect.ValueOf(v))
	})

	env.Define("toBool", func(v interface{}) bool {
		return util.ToBool(reflect.ValueOf(v))
	})

	env.Define("toChar", func(s rune) string {
		return string(s)
	})

	env.Define("toRune", func(s string) rune {
		if len(s) == 0 {
			return 0
		}
		return []rune(s)[0]
	})

	env.Define("toByteSlice", func(s string) []byte {
		return []byte(s)
	})

	env.Define("toRuneSlice", func(s string) []rune {
		return []rune(s)
	})

	env.Define("toBoolSlice", func(v []interface{}) []bool {
		var result []bool
		util.ToSlice(v, &result)
		return result
	})

	env.Define("toIntSlice", func(v []interface{}) []int64 {
		var result []int64
		util.ToSlice(v, &result)
		return result
	})

	env.Define("toFloatSlice", func(v []interface{}) []float64 {
		var result []float64
		util.ToSlice(v, &result)
		return result
	})

	env.Define("toStringSlice", func(v []interface{}) []string {
		var result []string
		util.ToSlice(v, &result)
		return result
	})

	env.Define("typeOf", func(v interface{}) string {
		t := reflect.TypeOf(v)
		if t == nil {
			return "<nil>"
		}
		return reflect.TypeOf(v).String()
	})

	env.Define("defined", func(s string) bool {
		_, err := env.Get(s)
		return err == nil
	})

	env.Define("print", fmt.Print)
	env.Define("println", fmt.Println)
	env.Define("printf", fmt.Printf)

	env.DefineType("int", int64(0))
	env.DefineType("float", float64(0.0))
	env.DefineType("bool", true)
	env.DefineType("string", "")
}
