package goblin

import (
	"errors"
	"fmt"
	"reflect"
)

// CallExpr defines a calling expression.
type CallExpr struct {
	PosImpl
	Func     interface{}
	Name     string
	SubExprs []Expr
	VarArg   bool
}

// Invoke the expression and return a result.
func (expr *CallExpr) Invoke(env *Env) (reflect.Value, error) {
	f := NilValue
	if expr.Func != nil {
		var ok bool
		if f, ok = expr.Func.(reflect.Value); !ok {
			f = NilValue
		}
	} else {
		var err error
		ff, err := env.Get(expr.Name)
		if err != nil {
			return f, err
		}
		f = ff
	}
	_, isReflect := f.Interface().(Func)

	args := []reflect.Value{}
	l := len(expr.SubExprs)
	for i, subExpr := range expr.SubExprs {
		arg, err := subExpr.Invoke(env)
		if err != nil {
			return arg, NewError(subExpr, err)
		}

		if i < f.Type().NumIn() {
			if !f.Type().IsVariadic() {
				it := f.Type().In(i)
				if arg.Kind().String() == "unsafe.Pointer" {
					arg = reflect.New(it).Elem()
				}
				if arg.Kind() != it.Kind() && arg.IsValid() && arg.Type().ConvertibleTo(it) {
					arg = arg.Convert(it)
				} else if arg.Kind() == reflect.Func {
					if _, isFunc := arg.Interface().(Func); isFunc {
						rfunc := arg
						arg = reflect.MakeFunc(it, func(args []reflect.Value) []reflect.Value {
							for i := range args {
								args[i] = reflect.ValueOf(args[i])
							}
							var rets []reflect.Value
							for _, v := range rfunc.Call(args)[:it.NumOut()] {
								rets = append(rets, v.Interface().(reflect.Value))
							}
							return rets
						})
					}
				} else if !arg.IsValid() {
					arg = reflect.Zero(it)
				}
			}
		}
		if !arg.IsValid() {
			arg = NilValue
		}

		if !isReflect {
			if expr.VarArg && i == l-1 {
				for j := 0; j < arg.Len(); j++ {
					args = append(args, arg.Index(j).Elem())
				}
			} else {
				args = append(args, arg)
			}
		} else {
			if arg.Kind() == reflect.Interface {
				arg = arg.Elem()
			}
			if expr.VarArg && i == l-1 {
				for j := 0; j < arg.Len(); j++ {
					args = append(args, reflect.ValueOf(arg.Index(j).Elem()))
				}
			} else {
				args = append(args, reflect.ValueOf(arg))
			}
		}
	}
	ret := NilValue
	var err error
	fnc := func() {
		defer func() {
			if ex := recover(); ex != nil {
				if e, ok := ex.(error); ok {
					err = e
				} else {
					err = errors.New(fmt.Sprint(ex))
				}
			}
		}()
		if f.Kind() == reflect.Interface {
			f = f.Elem()
		}
		rets := f.Call(args)
		if isReflect {
			var ok bool
			ev := rets[1].Interface()
			if ev != nil {
				if err, ok = ev.(error); !ok {
					err = nil
				}
			}
			if ret, ok = rets[0].Interface().(reflect.Value); !ok {
				ret = NilValue
			}
		} else {
			for i, subExpr := range expr.SubExprs {
				if ae, ok := subExpr.(*AddrExpr); ok {
					if id, ok := ae.Expr.(*IdentExpr); ok {
						_, err = id.Assign(args[i].Elem().Elem(), env)
					}
				}
			}
			if f.Type().NumOut() == 1 {
				ret = rets[0]
			} else {
				var result []interface{}
				for _, r := range rets {
					result = append(result, r.Interface())
				}
				ret = reflect.ValueOf(result)
			}
		}
	}
	fnc()
	if err != nil {
		return ret, NewError(expr, err)
	}
	return ret, nil
}

// Assign a value to the expression and return it.
func (expr *CallExpr) Assign(rv reflect.Value, env *Env) (reflect.Value, error) {
	return NilValue, newInvalidOperationError(expr)
}
