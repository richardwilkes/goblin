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
	result := []bool{false, true, true, true, false, false, false}
	require.Equal(t, len(input), len(result))
	for i := range input {
		assert.Equal(t, result[i], util.IsNumber(reflect.ValueOf(input[i])), "%v", input[i])
	}
}

func TestEqual(t *testing.T) {
	left := []interface{}{nil, "1", 2, 3.0, &[]string{}, &[]string{"hello"}, "goodbye"}
	right := []interface{}{nil, "1", 2, 2.0, 3, 3.0, &[]string{}, &[]string{"hello"}, "goodbye"}
	result := [][]bool{
		[]bool{true, false, false, false, false, false, false, false, false},
		[]bool{false, true, false, false, false, false, false, false, false},
		[]bool{false, false, true, true, false, false, false, false, false},
		[]bool{false, false, false, false, true, true, false, false, false},
		[]bool{false, false, false, false, false, false, true, false, false},
		[]bool{false, false, false, false, false, false, false, true, false},
		[]bool{false, false, false, false, false, false, false, false, true},
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
