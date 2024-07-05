package interp

import (
	"reflect"
	"testing"
)

func TestGenericFuncDeclare(t *testing.T) {
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

func TestGenericFuncComplex(t *testing.T) {
	i := New(Options{})
	done := false
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Do":    reflect.ValueOf(func() { done = true }),
			"Hello": reflect.ValueOf(GenericFunc("import . \"guthib.com/generic\"\nfunc Hello[T comparable, F any](v T, f func(a T) F) *T {\n\tDo(); return &v\n}")),
		},
	})
	i.ImportUsed()
	if err != nil {
		t.Error(err)
	}
	res, err := i.Eval("generic.Hello[int, bool](3, func(a int) bool { return true })")
	if err != nil {
		t.Error(err)
	}
	if res.Elem().Interface() != 3 {
		t.Error("expected &(3), got", res)
	}
	if !done {
		t.Error("!done")
	}
}

func TestCallPackageFunc(t *testing.T) {
	i := New(Options{})
	done := false
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Do": reflect.ValueOf(func() { done = true }),
		},
	})
	if err != nil {
		t.Error(err)
	}
	_, err = i.Eval("import . \"guthib.com/generic\"\nfunc Hello() { Do() }\nfunc main() { Hello() }")
	if err != nil {
		t.Error(err)
	}
	if !done {
		t.Error("!done")
	}
}
