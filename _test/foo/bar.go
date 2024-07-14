package foo

import "github.com/cogentcore/yaegi/_test/foo/boo"

var Bar = "BARR"
var Boo = boo.Boo

func init() { println("init foo") }
