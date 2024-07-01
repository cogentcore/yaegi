package interp

import (
	"fmt"
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
	fmt.Println(res)
}
