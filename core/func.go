package core

import (
	"github.com/spf13/cast"
	"reflect"
	"time"
)

func Equal(left, right interface{}) interface{} {
	return left == right
}

func covertFloat64(left, right interface{}) (float64, float64) {
	l := cast.ToFloat64(left)
	r := cast.ToFloat64(right)
	return l, r
}

func Plus(left, right interface{}) interface{} {
	l, r := covertFloat64(left, right)
	return l + r

}

func Sub(left, right interface{}) interface{} {
	leftKind := reflect.TypeOf(left)
	rightKind := reflect.TypeOf(right)

	if rightKind == reflect.TypeOf(time.Time{}) && leftKind == rightKind {
		// 处理日期相减的情况
		return int(left.(time.Time).Sub(right.(time.Time)).Hours() / 24)
	}

	l := cast.ToFloat64(left)
	r := cast.ToFloat64(right)
	return l - r
}

func Mul(left, right interface{}) interface{} {
	l := cast.ToFloat64(left)
	r := cast.ToFloat64(right)
	return l * r
}

func Div(left, right interface{}) interface{} {
	l := cast.ToFloat64(left)
	r := cast.ToFloat64(right)
	return l / r
}

func compare(left, right interface{}) int {
	leftKind := reflect.TypeOf(left)
	rightKind := reflect.TypeOf(right)

	if rightKind == reflect.TypeOf(time.Time{}) && leftKind == rightKind {
		return int(left.(time.Time).Sub(right.(time.Time)).Hours())
	}

	if ls, ok := left.(string); ok {
		if rs, ok := right.(string); ok {
			if ls == rs {
				return 0
			}
			if ls < rs {
				return -1
			}
			return 1
		}
	}
	if ls, ok := left.(float64); ok {
		if rs, ok := right.(float64); ok {
			return int(ls - rs)
		}
	}
	if ls, ok := left.(int64); ok {
		if rs, ok := right.(int64); ok {
			return int(ls - rs)
		}
	}
	if ls, ok := left.(uint64); ok {
		if rs, ok := right.(uint64); ok {
			return int(ls - rs)
		}
	}
	return 1
}

func Gt(left, right interface{}) interface{} {
	return compare(left, right) > 0
}

func Ge(left, right interface{}) interface{} {
	return compare(left, right) >= 0
}

func Lt(left, right interface{}) interface{} {
	return compare(left, right) < 0
}

func Le(left, right interface{}) interface{} {
	return compare(left, right) <= 0
}

func Ne(left, right interface{}) interface{} {
	return compare(left, right) != 0
}

func And(left, right interface{}) interface{} {
	return left.(bool) && right.(bool)
}

func Or(left, right interface{}) interface{} {
	return left.(bool) || right.(bool)
}

func ListContainsValue(val interface{}, array []interface{}, typ reflect.Type) bool {
	t := reflect.TypeOf(val)
	var nv interface{}
	if t.ConvertibleTo(typ) {
		nv = reflect.ValueOf(val).Convert(typ).Interface()
	}
	for _, x := range array {
		if x == nv {
			return true
		}
	}
	return false
}

func ListNotContainsValue(val interface{}, array []interface{}, tpe reflect.Type) bool {
	t := reflect.TypeOf(val)
	var nv interface{}
	if t.ConvertibleTo(tpe) {
		nv = reflect.ValueOf(val).Convert(tpe).Interface()
	}
	for _, x := range array {
		if x == nv {
			return false
		}
	}
	return true
}

func In(left, right interface{}) interface{} {
	if len(right.([]interface{})) <= 0 {
		return false
	}
	kind := reflect.TypeOf(right.([]interface{})[0])
	return ListContainsValue(left, right.([]interface{}), kind)
}

func NotIn(left, right interface{}) interface{} {
	if len(right.([]interface{})) <= 0 {
		return false
	}
	kind := reflect.TypeOf(right.([]interface{})[0])
	return ListNotContainsValue(left, right.([]interface{}), kind)
}
