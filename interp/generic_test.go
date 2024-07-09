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

func TestGenericFuncInfer(t *testing.T) {
	i := New(Options{})
	err := i.Use(Exports{
		"guthib.com/generic/generic": map[string]reflect.Value{
			"New":   reflect.ValueOf(GenericFunc("func New[T any]() *T { return new(T) }")),
			"AddAt": reflect.ValueOf(GenericFunc("func AddAt[T any](init func(n *T)) { v := New[T](); init(any(v).(*T)); println(*v) }")),
		},
	})
	i.ImportUsed()
	if err != nil {
		t.Error(err)
	}
	_, err = i.Eval(`
func main() {
	generic.AddAt(func(w *int) { *w = 3 })
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestGenericFuncInferDirect(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func New[T any]() *T { return new(T) }
func AddAt[T any](init func(n *T)) { v := New[T](); init(v); println(*v) }
func main() {
	AddAt(func(w *int) { *w = 3 })
}
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

func TestGenericFuncVarArg(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func New[T any](parent ...any) *T {
 	n := new(T)
	for _, p := range parent {
		println(p)
	}
}
func main() {
	New[int]()
	New[string]("happy")
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestForRangeIntBasic(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func main() {
	for i := range 3 {
		println(i)
	}
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestForRangeSlice(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func main() {
	s := []int{1,2,3}
	for i, v := range s {
		println(i,v)
	}
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestForRangeClosure(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func main() {
	fs := []func()
	for i := range 3 {
		println(i, &i)
		fs = append(fs, func() { println(i, &i) })
	}
	for _, f := range fs {
		f()
	}
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestForLoopClosure(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func main() {
	fs := []func()
	for i := 0; i < 3; i++ {
		println(i)
		f := func() { println(i) }
		fs = append(fs, f)
	}
	for _, f := range fs {
		f()
	}
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestForLoopAddr(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func main() {
	is := []*int
	for i := 0; i < 3; i++ {
		is = append(is, &i)
		println(i)
	}
	for _, j := range is {
		println(*j)
	}
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestForRangeAddr(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func main() {
	is := []*int
	for i := range 3 {
		println(i)
		is = append(is, &i)
	}
	for j := 0; j < len(is); j++ {
		println(*is[j])
	}
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestClosureFailure(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func main() {
	fs := []func()
	add := func(fun func()) {
		fs = append(fs, fun)
	}
	
	for i := range 3 {
		println(i)
		add(func() { println(i) })
	}
	for _, f := range fs {
		f()
	}
}
`)
	if err != nil {
		t.Error(err)
	}
}
