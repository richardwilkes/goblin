package goblin

import "reflect"

func isNil(v reflect.Value) bool {
	if !v.IsValid() || v.Kind().String() == "unsafe.Pointer" {
		return true
	}
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.IsNil() {
		return true
	}
	return false
}

func isNum(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func equal(left, right reflect.Value) bool {
	leftIsNil, rightIsNil := isNil(left), isNil(right)
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
	if isNum(left) && isNum(right) {
		if right.Type().ConvertibleTo(left.Type()) {
			right = right.Convert(left.Type())
		}
	}
	if left.CanInterface() && right.CanInterface() {
		return reflect.DeepEqual(left.Interface(), right.Interface())
	}
	return reflect.DeepEqual(left, right)
}
