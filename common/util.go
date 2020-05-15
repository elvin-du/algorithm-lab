package common

import "reflect"

func IsNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}
