package interp

import (
	"testing"
)

func TestGenericRecursive(t *testing.T) {
	i := New(Options{})
	_, err := i.Eval(`
func New[T any]() *T { return new(T) }
func Add[T any]() *T { return New[T]() }

func main() {
	v := Add[int]() // crashes!
//	v := New[int]() // direct calling works
	*v = 3
	println(*v)
}
`)
	if err != nil {
		t.Error(err)
	}
}

func TestGenericRecursiveFunc(t *testing.T) {
	t.Skip("not yet")
	i := New(Options{})
	_, err := i.Eval(`
func New[T any]() *T { return new(T) }
func Add[T any](init func(n *T)) { v := New(T); init(v); println(*v) }

func main() {
// v := Add[int]() // crashes!
// v := Add() // not enough info for non-init version
Add(func(w *int) { *w = 3 }) // for init version
// v := New[int]() // direct calling works
}
`)
	if err != nil {
		t.Error(err)
	}
}
