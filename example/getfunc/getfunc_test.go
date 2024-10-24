package main

import (
	"fmt"
	"testing"

	"github.com/breadchris/yaegi/interp"
	"github.com/breadchris/yaegi/stdlib"
)

func TestGetFunc(t *testing.T) {
	i := interp.New(interp.Options{GoPath: "./_gopath/"})
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(`import "github.com/foo/bar"`); err != nil {
		t.Fatal(err)
	}

	val, err := i.Eval(`bar.NewFoo`)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(val.Call(nil))
}
