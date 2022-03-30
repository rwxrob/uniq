/*
Package foo is where the stand-alone high-level functions go that Bonzai
branch depends on. This way people can import just pkg and use the
functions directly from anything, not just Bonzai branches an trees. It
is kept in pkg to avoid cyclical import problems.
*/
package foo

import "fmt"

func Foo() { fmt.Println("foo") }
