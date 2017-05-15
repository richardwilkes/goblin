package goblin

import (
	"math"
	"reflect"
	"strings"
)

// BinOpExpr defines a binary operator expression.
type BinOpExpr struct {
	PosImpl
	LHS      Expr
	Operator string
	RHS      Expr
}

// Invoke the expression and return a result.
func (expr *BinOpExpr) Invoke(env *Env) (reflect.Value, error) {
	LHSV := NilValue
	RHSV := NilValue
	var err error

	LHSV, err = expr.LHS.Invoke(env)
	if err != nil {
		return LHSV, NewError(expr, err)
	}
	if LHSV.Kind() == reflect.Interface {
		LHSV = LHSV.Elem()
	}
	if expr.RHS != nil {
		RHSV, err = expr.RHS.Invoke(env)
		if err != nil {
			return RHSV, NewError(expr, err)
		}
		if RHSV.Kind() == reflect.Interface {
			RHSV = RHSV.Elem()
		}
	}
	switch expr.Operator {
	case "+":
		if LHSV.Kind() == reflect.String || RHSV.Kind() == reflect.String {
			return reflect.ValueOf(toString(LHSV) + toString(RHSV)), nil
		}
		if (LHSV.Kind() == reflect.Array || LHSV.Kind() == reflect.Slice) && (RHSV.Kind() != reflect.Array && RHSV.Kind() != reflect.Slice) {
			return reflect.Append(LHSV, RHSV), nil
		}
		if (LHSV.Kind() == reflect.Array || LHSV.Kind() == reflect.Slice) && (RHSV.Kind() == reflect.Array || RHSV.Kind() == reflect.Slice) {
			return reflect.AppendSlice(LHSV, RHSV), nil
		}
		if LHSV.Kind() == reflect.Float64 || RHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(toFloat64(LHSV) + toFloat64(RHSV)), nil
		}
		return reflect.ValueOf(toInt64(LHSV) + toInt64(RHSV)), nil
	case "-":
		if LHSV.Kind() == reflect.Float64 || RHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(toFloat64(LHSV) - toFloat64(RHSV)), nil
		}
		return reflect.ValueOf(toInt64(LHSV) - toInt64(RHSV)), nil
	case "*":
		if LHSV.Kind() == reflect.String && (RHSV.Kind() == reflect.Int || RHSV.Kind() == reflect.Int32 || RHSV.Kind() == reflect.Int64) {
			return reflect.ValueOf(strings.Repeat(toString(LHSV), int(toInt64(RHSV)))), nil
		}
		if LHSV.Kind() == reflect.Float64 || RHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(toFloat64(LHSV) * toFloat64(RHSV)), nil
		}
		return reflect.ValueOf(toInt64(LHSV) * toInt64(RHSV)), nil
	case "/":
		return reflect.ValueOf(toFloat64(LHSV) / toFloat64(RHSV)), nil
	case "%":
		return reflect.ValueOf(toInt64(LHSV) % toInt64(RHSV)), nil
	case "==":
		return reflect.ValueOf(equal(LHSV, RHSV)), nil
	case "!=":
		return reflect.ValueOf(!equal(LHSV, RHSV)), nil
	case ">":
		return reflect.ValueOf(toFloat64(LHSV) > toFloat64(RHSV)), nil
	case ">=":
		return reflect.ValueOf(toFloat64(LHSV) >= toFloat64(RHSV)), nil
	case "<":
		return reflect.ValueOf(toFloat64(LHSV) < toFloat64(RHSV)), nil
	case "<=":
		return reflect.ValueOf(toFloat64(LHSV) <= toFloat64(RHSV)), nil
	case "|":
		return reflect.ValueOf(toInt64(LHSV) | toInt64(RHSV)), nil
	case "||":
		if toBool(LHSV) {
			return LHSV, nil
		}
		return RHSV, nil
	case "&":
		return reflect.ValueOf(toInt64(LHSV) & toInt64(RHSV)), nil
	case "&&":
		if toBool(LHSV) {
			return RHSV, nil
		}
		return LHSV, nil
	case "**":
		if LHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(math.Pow(toFloat64(LHSV), toFloat64(RHSV))), nil
		}
		return reflect.ValueOf(int64(math.Pow(toFloat64(LHSV), toFloat64(RHSV)))), nil
	case ">>":
		return reflect.ValueOf(toInt64(LHSV) >> uint64(toInt64(RHSV))), nil
	case "<<":
		return reflect.ValueOf(toInt64(LHSV) << uint64(toInt64(RHSV))), nil
	default:
		return NilValue, NewStringError(expr, "Unknown operator")
	}
}

// Assign a value to the expression and return it.
func (expr *BinOpExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
