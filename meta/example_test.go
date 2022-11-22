// meta package provides abilities to meta-programming, and makes reflect package make sense.
// Functions in this package MUST NOT panic.
package meta

import (
	"fmt"
	"runtime"
)

func ExampleIsUntypedNil() {
	fmt.Println(IsUntypedNil(nil)) // true

	var p1 *int
	fmt.Println(IsUntypedNil(p1)) // false

	var p2 error                  // error is an interface
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

	var p2 error                // error is an interface
	fmt.Println(IsTypedNil(p2)) // false

	// Output:
	// false
	// true
	// false
}

func ExampleIsNil_why() {
	// Go sucks. It's not possible to compare a typed nil to nil.
	var err error
	fmt.Printf("(nil == nil) -> %v\n", err == nil) // true
	fmt.Printf("type = %T\n", err)                 // <nil>
	// err is an untyped nil now.
	u1, t1 := IsNil(err)
	fmt.Printf("IsNil(err) -> %v, %v\n", u1, t1) // true, false

	err = func() error {
		// some function returns an error but you don't know what type it actually is.
		return func() error {
			// n is a typed nil
			var n *runtime.TypeAssertionError = nil
			return n
		}()
	}()
	fmt.Printf("(nil == nil) -> %v\n", err == nil) // false
	fmt.Printf("type = %T\n", err)                 // *runtime.TypeAssertionError
	// err is a typed nil now.
	u2, t2 := IsNil(err)
	fmt.Printf("IsNil(err) -> %v, %v\n", u2, t2) // false, true

	// Output:
	// (nil == nil) -> true
	// type = <nil>
	// IsNil(err) -> true, false
	// (nil == nil) -> false
	// type = *runtime.TypeAssertionError
	// IsNil(err) -> false, true
}

func ExampleIsNil_how() {
	var err error

	if err == nil {
		fmt.Println("err == nil")
	}

	if u, t := IsNil(err); u || t {
		fmt.Println("err is nil")
	}

	err = func() error {
		// some function returns an error but you don't know what type it actually is.
		return func() error {
			// n is a typed nil
			var n *runtime.TypeAssertionError = nil
			return n
		}()
	}()

	if err != nil {
		fmt.Println("err != nil")
	}

	if u, t := IsNil(err); u || t {
		fmt.Println("err is nil")
	}

	// Output:
	// err == nil
	// err is nil
	// err != nil
	// err is nil
}
