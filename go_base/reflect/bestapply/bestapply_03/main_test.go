package main

import (
	"reflect"
	"testing"
)

func TestReflectFunc(t *testing.T) {
	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}
	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}
	var (
		function reflect.Value
		invalue  []reflect.Value
		n        int
	)
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		invalue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			invalue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(invalue)
	}
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}
