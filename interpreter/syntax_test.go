package interpreter_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/richardwilkes/goblin"
	"github.com/richardwilkes/goblin/interpreter/parser"
	"github.com/richardwilkes/goblin/util"
	"github.com/stretchr/testify/assert"
)

func TestComments(t *testing.T) {
	checkSyntax(t, "# single line comment")
	checkSyntax(t, "x = 1 # single line comment")
	checkSyntax(t, "x = 1# single line comment")
	checkSyntax(t, "// single line comment")
	checkSyntax(t, "x = 1 // single line comment")
	checkSyntax(t, "x = 1// single line comment")
	checkSyntax(t, `/*
 * multi-line comment
 */`)
	checkSyntax(t, `x = 1/*
 * multi-line comment
 */
y = 2`)
	checkSyntax(t, `for i = 0; i < 10 /* comment */; i++ {}`)
}

func checkSyntax(t *testing.T, script string) {
	_, err := parser.Parse(script)
	assert.NoError(t, err, script)
}

func TestVariableDeclaration(t *testing.T) {
	checkDeclaration(t, "x = 1", 1)
	checkDeclaration(t, "x = 1.2", 1.2)
	checkDeclaration(t, "x = 1.2e3", 1.2e3)
	checkDeclaration(t, "x = -1", -1)
	checkDeclaration(t, "x = -1.2", -1.2)
	checkDeclaration(t, "x = -1.2e3", -1.2e3)
	checkDeclaration(t, "x = true", true)
	checkDeclaration(t, "x = false", false)
	checkDeclaration(t, "x = [ 1, 2, 3 ]", []interface{}{int64(1), int64(2), int64(3)})
	checkDeclaration(t, `x = { "foo": "bar", "bar": "baz" }`, map[string]interface{}{"foo": "bar", "bar": "baz"})
	checkDeclaration(t, `x = {
"foo": "bar",
"bar": {
	"one": 1,
	"two": true,
	"three": [ false, 1.23e4, "hello" ]},
}`,
		map[string]interface{}{
			"foo": "bar",
			"bar": map[string]interface{}{
				"one":   int64(1),
				"two":   true,
				"three": []interface{}{false, 1.23e4, "hello"},
			},
		})
	checkDeclaration(t, `x = "hello"`, "hello")
	checkDeclaration(t, `x = "1"`, "1")
	checkDeclaration(t, `x = ""`, "")
	checkDeclaration(t, `x = nil`, nil)
}

func TestVariableConversion(t *testing.T) {
	data := []string{"1", "1.2", "1.2e3", "true", "false", "[1,2,3]", `{"foo":1}`, `"foo"`, "nil"}
	types := []string{"int64", "float64", "float64", "bool", "bool", "[]interface {}", "map[string]interface {}", "string", "<nil>"}
	for i := range data {
		checkDeclaration(t, fmt.Sprintf("typeOf(%s)", data[i]), types[i])
	}
	stringResults := []string{"1", "1.2", "1200", "true", "false", "[1 2 3]", "map[foo:1]", "foo", "<nil>"}
	for i := range data {
		checkDeclaration(t, fmt.Sprintf("toString(%s)", data[i]), stringResults[i])
	}
	intResults := []int64{1, 1, 1200, 1, 0, 0, 0, 0, 0}
	for i := range data {
		checkDeclaration(t, fmt.Sprintf("toInt(%s)", data[i]), intResults[i])
	}
	floatResults := []float64{1, 1.2, 1200, 1, 0, 0, 0, 0, 0}
	for i := range data {
		checkDeclaration(t, fmt.Sprintf("toFloat(%s)", data[i]), floatResults[i])
	}
	boolResults := []bool{true, true, true, true, false, false, false, false, false}
	for i := range data {
		checkDeclaration(t, fmt.Sprintf("toBool(%s)", data[i]), boolResults[i])
	}
	for i := 32; i <= 1000; i++ {
		if i != '"' && i != '\\' {
			checkDeclaration(t, fmt.Sprintf(`toRune("%c")`, rune(i)), rune(i))
		}
	}
	for i := 32; i <= 1000; i++ {
		if i != '"' && i != '\\' {
			checkDeclaration(t, fmt.Sprintf(`toChar(%d)`, i), fmt.Sprintf("%c", rune(i)))
		}
	}
	checkDeclaration(t, `toByteSlice("foo")`, []byte("foo"))
	checkDeclaration(t, `toRuneSlice("foo")`, []rune("foo"))
	checkDeclaration(t, `toBoolSlice([true, false])`, []bool{true, false})
	checkDeclaration(t, `toIntSlice([1,2])`, []int64{1, 2})
	checkDeclaration(t, `toFloatSlice([1.2, 3.4])`, []float64{1.2, 3.4})
	checkDeclaration(t, `toStringSlice(["a", "b"])`, []string{"a", "b"})
	checkDeclaration(t, `defined("x")`, false)
	checkDeclaration(t, `x = 1; defined("x")`, true)
}

func TestRangeLenKeys(t *testing.T) {
	checkDeclaration(t, `len("hello")`, 5)
	checkDeclaration(t, `len("")`, 0)
	checkDeclaration(t, "len([])", 0)
	checkDeclaration(t, "len([1,2,3,4])", 4)
	checkDeclaration(t, "len({})", 0)
	checkDeclaration(t, `len({"foo":1, "bar":2})`, 2)
	checkDeclaration(t, `v = 0; for i in [1,2,3] { v += i }; v`, 6)
	checkDeclaration(t, `v = 0; a = [1,2,3]; for i in a { v += i }; v`, 6)
	checkDeclaration(t, `v = 0; for i in range(3,6) { v += i }; v`, 18)

	script := `x = {"foo":1, "bar":2}; keys(x)`
	v, err := goblin.ParseAndRun(script)
	if assert.NoError(t, err, script) {
		if assert.Equal(t, reflect.Slice, v.Kind()) {
			if assert.Equal(t, 2, v.Len()) {
				v1 := v.Index(0).Interface()
				v2 := v.Index(1).Interface()
				assert.True(t, (v1 == "foo" && v2 == "bar") || (v1 == "bar" && v2 == "foo"), `["%v", "%v"]`, v1, v2)
			}
		}
	}
}

func TestOperators(t *testing.T) {
	checkDeclaration(t, "1 > 0", true)
	checkDeclaration(t, "1 <= 1", true)
	checkDeclaration(t, "1.0 <= 1.0", true)
	checkDeclaration(t, "1 == 1.0", true)
	checkDeclaration(t, `1 != "1"`, false)
	checkDeclaration(t, "1 == 1", true)
	checkDeclaration(t, "1.2 == 1.2", true)
	checkDeclaration(t, `"1" == "1"`, true)
	checkDeclaration(t, `false != "1"`, true)
	checkDeclaration(t, "false != true", true)
	checkDeclaration(t, "true == true", true)
	checkDeclaration(t, "nil == nil", true)
	checkDeclaration(t, `1 <= 2 ? "a" : "b"`, "a")
	checkDeclaration(t, `2 <= 1 ? "a" : "b"`, "b")
	checkDeclaration(t, "a = 1; a += 2", 3)
	checkDeclaration(t, "a = 1; a -= 2", -1)
	checkDeclaration(t, "a = 10; a--", 9)
	checkDeclaration(t, "a = 10; a++", 11)
	checkDeclaration(t, "a = 10; a *= 2", 20)
	checkDeclaration(t, "a = 10; a /= 2", 5)
	checkDeclaration(t, "a = 2**3", 8)
	checkDeclaration(t, "a = 1; a &= 2", 0)
	checkDeclaration(t, "a = 3; a &= 2", 2)
	checkDeclaration(t, "a = 1; a |= 2", 3)
	checkDeclaration(t, "a = !3", false)
	checkDeclaration(t, "a = !!3", true)
	checkDeclaration(t, "a = !true", false)
	checkDeclaration(t, "a = !false", true)
	checkDeclaration(t, "a = ^3", -4)
	checkDeclaration(t, "a = 3 << 2", 12)
	checkDeclaration(t, "a = 11 >> 2", 2)
}

func TestFunc(t *testing.T) {
	checkDeclaration(t, "x = func a() { }()", nil)
	checkDeclaration(t, "func a() { return 2 }; a()", 2)
	checkDeclaration(t, "func a(x) { return x + 1 }; a(5)", 6)
	checkDeclaration(t, "func a(x) { return x + 1, x + 2 }; a(5)", []interface{}{int64(6), int64(7)})
	checkDeclaration(t, `
x = func(x) {
  return func(y) {
    x(y)
  }
}(func(z) {
  return "Yay! " + z
})("hello world")
	`, "Yay! hello world")
}

func TestFor(t *testing.T) {
	checkDeclaration(t, `
x = 0
for i in [2,4,6] {
	x += i
}
x	`, 12)
	checkDeclaration(t, `
x = 0
for {
	x++
	if x > 3 {
		break
	}
}
x	`, 4)
	checkDeclaration(t, `
func loop() {
	x = 0
	for {
		if x == 5 {
			return x
		}
		x++
	}
	return 1
}
loop()`, 5)
	checkDeclaration(t, `
func loop() {
	x = 0
	for i in range(0,10) {
		if i == 5 {
			return x
		}
		x++
	}
	return 1
}
loop()`, 5)
	checkDeclaration(t, `
func loop() {
	x = 0
	for i = 0; i < 10; i++ {
		if i == 5 {
			return x
		}
		x++
	}
	return 1
}
loop()`, 5)
	checkDeclaration(t, `
r = {
	"stuff": [
		{
			"x": 1,
			"y": 2,
		},
		{
			"x": 5,
			"y": -2,
		},
	],
}
x = 0
for i in r.stuff {
	x += i.x
}
x`, 6)
}

func TestSwitch(t *testing.T) {
	checkDeclaration(t, `
x = 0
r = -1
switch x {
	case 0:
		r = 1
	case 1:
		r = 3
	default:
		r = 6
}
r`, 1)
	checkDeclaration(t, `
x = 1
r = -1
switch x {
	case 0:
		r = 1
	case 1:
		r = 3
	default:
		r = 6
}
r`, 3)
	checkDeclaration(t, `
x = 2
r = -1
switch x {
	case 0:
		r = 1
	case 1:
		r = 3
	default:
		r = 6
}
r`, 6)
	checkDeclaration(t, `
x = 2
r = -1
switch x {
	case 0:
		r = 1
	case 1:
		r = 3
}
r`, -1)
}

func TestIf(t *testing.T) {
	checkDeclaration(t, `
x = -1
if true {
	x = 1
} else if true {
	x = 2
} else {
	x = 3
}
x`, 1)
	checkDeclaration(t, `
x = -1
if false {
	x = 1
} else if true {
	x = 2
} else {
	x = 3
}
x`, 2)
	checkDeclaration(t, `
x = -1
if false {
	x = 1
} else if false {
	x = 2
} else {
	x = 3
}
x`, 3)
}

func TestSort(t *testing.T) {
	checkDeclaration(t, `
a = [3,1,2]
sort(a, func(i, j) { return a[i] < a[j] })
a
`, []interface{}{int64(1), int64(2), int64(3)})
}

func checkDeclaration(t *testing.T, script string, expected interface{}) {
	v, err := goblin.ParseAndRun(script)
	if assert.NoError(t, err, script) {
		assert.True(t, util.Equal(reflect.ValueOf(expected), v), "script: %s\nexpected: %v\nreceived: %v", script, expected, v)
	}
}
