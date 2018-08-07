package mock

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//func A() string {
//	return "hello, world"
//}

func TestMock(t *testing.T) {
	A := func() string {
		return "hello, world"
	}
	assert.Equal(t, "hello, world", A())

	var ptr uintptr = reflect.ValueOf(A).Pointer()
	replace(ptr)
	//replace(&A)

	assert.Equal(t, "hello, A", A())
}

func replace(a interface{}) {
	fn := reflect.ValueOf(a).Elem()
	mockFn := reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
		v := reflect.ValueOf("hello, A")
		return []reflect.Value{v}
	})
	fn.Set(mockFn)
}
