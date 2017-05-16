package expression

import (
	"math"
	"reflect"
	"strings"

	"github.com/richardwilkes/goblin/interpreter"
	"github.com/richardwilkes/goblin/util"
)

// BinOp defines a binary operator expression.
type BinOp struct {
	interpreter.PosImpl
	LHS      interpreter.Expr
	Operator string
	RHS      interpreter.Expr
}

// Invoke the expression and return a result.
func (expr *BinOp) Invoke(env *interpreter.Env) (reflect.Value, error) {
	LHSV := interpreter.NilValue
	RHSV := interpreter.NilValue
	var err error

	LHSV, err = expr.LHS.Invoke(env)
	if err != nil {
		return LHSV, interpreter.NewError(expr, err)
	}
	if LHSV.Kind() == reflect.Interface {
		LHSV = LHSV.Elem()
	}
	if expr.RHS != nil {
		RHSV, err = expr.RHS.Invoke(env)
		if err != nil {
			return RHSV, interpreter.NewError(expr, err)
		}
		if RHSV.Kind() == reflect.Interface {
			RHSV = RHSV.Elem()
		}
	}
	switch expr.Operator {
	case "+":
		if LHSV.Kind() == reflect.String || RHSV.Kind() == reflect.String {
			return reflect.ValueOf(util.ToString(LHSV) + util.ToString(RHSV)), nil
		}
		if (LHSV.Kind() == reflect.Array || LHSV.Kind() == reflect.Slice) && (RHSV.Kind() != reflect.Array && RHSV.Kind() != reflect.Slice) {
			return reflect.Append(LHSV, RHSV), nil
		}
		if (LHSV.Kind() == reflect.Array || LHSV.Kind() == reflect.Slice) && (RHSV.Kind() == reflect.Array || RHSV.Kind() == reflect.Slice) {
			return reflect.AppendSlice(LHSV, RHSV), nil
		}
		if LHSV.Kind() == reflect.Float64 || RHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(util.ToFloat64(LHSV) + util.ToFloat64(RHSV)), nil
		}
		return reflect.ValueOf(util.ToInt64(LHSV) + util.ToInt64(RHSV)), nil
	case "-":
		if LHSV.Kind() == reflect.Float64 || RHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(util.ToFloat64(LHSV) - util.ToFloat64(RHSV)), nil
		}
		return reflect.ValueOf(util.ToInt64(LHSV) - util.ToInt64(RHSV)), nil
	case "*":
		if LHSV.Kind() == reflect.String && (RHSV.Kind() == reflect.Int || RHSV.Kind() == reflect.Int32 || RHSV.Kind() == reflect.Int64) {
			return reflect.ValueOf(strings.Repeat(util.ToString(LHSV), int(util.ToInt64(RHSV)))), nil
		}
		if LHSV.Kind() == reflect.Float64 || RHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(util.ToFloat64(LHSV) * util.ToFloat64(RHSV)), nil
		}
		return reflect.ValueOf(util.ToInt64(LHSV) * util.ToInt64(RHSV)), nil
	case "/":
		return reflect.ValueOf(util.ToFloat64(LHSV) / util.ToFloat64(RHSV)), nil
	case "%":
		return reflect.ValueOf(util.ToInt64(LHSV) % util.ToInt64(RHSV)), nil
	case "==":
		return reflect.ValueOf(util.Equal(LHSV, RHSV)), nil
	case "!=":
		return reflect.ValueOf(!util.Equal(LHSV, RHSV)), nil
	case ">":
		return reflect.ValueOf(util.ToFloat64(LHSV) > util.ToFloat64(RHSV)), nil
	case ">=":
		return reflect.ValueOf(util.ToFloat64(LHSV) >= util.ToFloat64(RHSV)), nil
	case "<":
		return reflect.ValueOf(util.ToFloat64(LHSV) < util.ToFloat64(RHSV)), nil
	case "<=":
		return reflect.ValueOf(util.ToFloat64(LHSV) <= util.ToFloat64(RHSV)), nil
	case "|":
		return reflect.ValueOf(util.ToInt64(LHSV) | util.ToInt64(RHSV)), nil
	case "||":
		if util.ToBool(LHSV) {
			return LHSV, nil
		}
		return RHSV, nil
	case "&":
		return reflect.ValueOf(util.ToInt64(LHSV) & util.ToInt64(RHSV)), nil
	case "&&":
		if util.ToBool(LHSV) {
			return RHSV, nil
		}
		return LHSV, nil
	case "**":
		if LHSV.Kind() == reflect.Float64 {
			return reflect.ValueOf(math.Pow(util.ToFloat64(LHSV), util.ToFloat64(RHSV))), nil
		}
		return reflect.ValueOf(int64(math.Pow(util.ToFloat64(LHSV), util.ToFloat64(RHSV)))), nil
	case ">>":
		return reflect.ValueOf(util.ToInt64(LHSV) >> uint64(util.ToInt64(RHSV))), nil
	case "<<":
		return reflect.ValueOf(util.ToInt64(LHSV) << uint64(util.ToInt64(RHSV))), nil
	default:
		return interpreter.NilValue, interpreter.NewStringError(expr, "Unknown operator")
	}
}

// Assign a value to the expression and return it.
func (expr *BinOp) Assign(rv reflect.Value, env *interpreter.Env) (reflect.Value, error) {
	return interpreter.NilValue, interpreter.NewInvalidOperationError(expr)
}
