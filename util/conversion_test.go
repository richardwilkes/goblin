package util_test

import (
	"reflect"
	"testing"

	"github.com/richardwilkes/goblin/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToString(t *testing.T) {
	input := []interface{}{"1", "1g", 2, 3.0, 4.5, true, false, nil}
	result := []string{"1", "1g", "2", "3", "4.5", "true", "false", "nil"}
	require.Equal(t, len(input), len(result))
	for i := range input {
		assert.Equal(t, result[i], util.ToString(reflect.ValueOf(input[i])), "%v", input[i])
	}
}

func TestToBool(t *testing.T) {
	input := []interface{}{-1, 0, 1, 2, 22.0, 0.0, "1", "0", "yes", "true", "TruE", "no", "y", "Y", "22", true, false, "0x0", "0x1", "0x10", "hello"}
	result := []bool{true, false, true, true, true, false, true, false, true, true, true, false, true, true, true, true, false, false, true, true, false}
	require.Equal(t, len(input), len(result))
	for i := range input {
		assert.Equal(t, result[i], util.ToBool(reflect.ValueOf(input[i])), "%v", input[i])
	}
}

func TestToInt64(t *testing.T) {
	input := []interface{}{21, 22.0, "23", "24.5", true, false, "0x10", "hello"}
	result := []int64{21, 22, 23, 24, 1, 0, 16, 0}
	require.Equal(t, len(input), len(result))
	for i := range input {
		assert.Equal(t, result[i], util.ToInt64(reflect.ValueOf(input[i])), "%v", input[i])
	}
}

func TestToFloat64(t *testing.T) {
	input := []interface{}{21, 22.0, "23", "24.5", true, false, "0x10", "hello"}
	result := []float64{21, 22, 23, 24.5, 1, 0, 16, 0}
	require.Equal(t, len(input), len(result))
	for i := range input {
		assert.Equal(t, result[i], util.ToFloat64(reflect.ValueOf(input[i])), "%v", input[i])
	}
}

func TestToIntSlice(t *testing.T) {
	generic := []interface{}{1, 2, 3, 4, 5}
	specific := []int{1, 2, 3, 4, 5}
	assert.NotEqual(t, specific, generic)
	var converted []int
	util.ToSlice(generic, &converted)
	assert.Equal(t, specific, converted)
}

func TestToFloat64Slice(t *testing.T) {
	generic := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
	specific := []float64{1, 2, 3, 4, 5}
	assert.NotEqual(t, specific, generic)
	var converted []float64
	util.ToSlice(generic, &converted)
	assert.Equal(t, specific, converted)
}

func TestToStringSlice(t *testing.T) {
	generic := []interface{}{"a", "b", "c"}
	specific := []string{"a", "b", "c"}
	assert.NotEqual(t, specific, generic)
	var converted []string
	util.ToSlice(generic, &converted)
	assert.Equal(t, specific, converted)
}
