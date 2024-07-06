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

func TestGenericFuncBasic(t *testing.T) {
	i := New(Options{})
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Hello": reflect.ValueOf(GenericFunc("func Hello[T comparable](v T) *T {\n\treturn &v\n}")),
		},
	})
	if err != nil {
		t.Error(err)
	}
	res, err := i.Eval("generic.Hello(3)")
	if err != nil {
		t.Error(err)
	}
	if res.Elem().Interface() != 3 {
		t.Error("expected &(3), got", res)
	}
}

func TestGenericFuncNoDotImport(t *testing.T) {
	i := New(Options{})
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Hello": reflect.ValueOf(GenericFunc("func Hello[T any](v T) { println(v) }")),
		},
	})
	if err != nil {
		t.Error(err)
	}
	_, err = i.Eval(`
import "guthib.com/generic"
func main() { generic.Hello(3) }
`)
	if err != nil {
		t.Error(err)
	}
}

func TestGenericFuncDotImport(t *testing.T) {
	i := New(Options{})
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Hello": reflect.ValueOf(GenericFunc("func Hello[T any](v T) { println(v) }")),
		},
	})
	if err != nil {
		t.Error(err)
	}
	_, err = i.Eval(`
import . "guthib.com/generic"
func main() { Hello(3) }
`)
	if err != nil {
		t.Error(err)
	}
}

func TestGenericFuncComplex(t *testing.T) {
	i := New(Options{})
	done := false
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Do":    reflect.ValueOf(func() { done = true }),
			"Hello": reflect.ValueOf(GenericFunc("func Hello[T comparable, F any](v T, f func(a T) F) *T {\n\tDo(); return &v\n}")),
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

func TestGenericCallTwice(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func Do[T any](v T) { println(v) }
func Hello[T any](v T) { Do(v) }
func main() { Hello[int](3) }
`)
	if err != nil {
		t.Error(err)
	}
}

func TestGenericFuncTwice(t *testing.T) {
	i := New(Options{})
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"Do":    reflect.ValueOf(GenericFunc("func Do[T any](v T) { println(v) }")),
			"Hello": reflect.ValueOf(GenericFunc("func Hello[T any](v T) { Do(v) }")),
		},
	})
	i.ImportUsed()
	if err != nil {
		t.Error(err)
	}
	_, err = i.Eval(`
func main() { generic.Hello[int](3) }
`)
	if err != nil {
		t.Error(err)
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
