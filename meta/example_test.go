// meta package provides abilities to meta-programming, and makes reflect package make sense.
package meta

import (
	"fmt"
	"io"
)

func ExampleIsUntypedNil() {
	fmt.Println(IsUntypedNil(nil)) // true

	var p1 *int
	fmt.Println(IsUntypedNil(p1)) // false

	var p2 io.Writer
	fmt.Println(IsUntypedNil(p2)) // true

	// Output:
	// true
	// false
	// true
}

func ExampleIsTypedNil() {
	fmt.Println(IsTypedNil(nil)) // false

	var p1 *int
	fmt.Println(IsTypedNil(p1)) // true

	var p2 io.Writer
	fmt.Println(IsTypedNil(p2)) // false

	// Output:
	// false
	// true
	// false
}
