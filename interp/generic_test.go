package interp

import (
	"reflect"
	"testing"
)

func TestGenericFunc(t *testing.T) {
	i := New(Options{})
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Hello": reflect.ValueOf(GenericFunc("func Hello[T comparable](v T) *T {\n\treturn &v\n}")),
		},
	})
	if err != nil {
		t.Error(err)
	}
	i.ImportUsed()
	res, err := i.Eval("generic.Hello(3)")
	if err != nil {
		t.Error(err)
	}
	if res.Elem().Interface() != 3 {
		t.Error("expected &(3), got", res)
	}
}

func TestDeclareGenericFunc(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval("func Hello[T comparable](v T) *T {\n\treturn &v\n}")
	if err != nil {
		t.Error(err)
	}
	res, err := i.Eval("Hello(3)")
	if err != nil {
		t.Error(err)
	}
	if res.Elem().Interface() != 3 {
		t.Error("expected &(3), got", res)
	}
}
