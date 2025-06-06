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

func ExampleIsNil_typed() {
	isUntyped, isTyped := false, false

	// nil to an interface type is untyped nil.
	var e1 error = nil
	isUntyped, isTyped = IsNil(e1)
	fmt.Printf("e1 == nil: %v\n", e1 == nil)                              // e1 == nil: true
	fmt.Printf("IsNil(e1) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(e1) -> untyped=true, typed=false

	// nil to a concrete type is also untyped nil, but it is typed with reflect.
	var n1 *runtime.TypeAssertionError
	isUntyped, isTyped = IsNil(n1)
	fmt.Printf("n1 == nil: %v\n", n1 == nil)                              // n1 == nil: true
	fmt.Printf("IsNil(n1) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(n1) -> untyped=false, typed=true

	// literal untyped nil directly returns from a function to interface type is untyped nil.
	var e2 error = func() error {
		return nil
	}()
	isUntyped, isTyped = IsNil(e2)
	fmt.Printf("e2 == nil: %v\n", e2 == nil)                              // e2 == nil: true
	fmt.Printf("IsNil(e2) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(e2) -> untyped=true, typed=false

	// literal untyped nil directly returns from a function to a concrete type is typed nil.
	var n2 *runtime.TypeAssertionError = func() *runtime.TypeAssertionError {
		return nil
	}()
	isUntyped, isTyped = IsNil(n2)
	fmt.Printf("n2 == nil: %v\n", n2 == nil)                              // n2 == nil: true
	fmt.Printf("IsNil(n2) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(n2) -> untyped=false, typed=true

	// literal untyped nil to interface directly returns from a function to interface type is untyped nil.
	var e3i error = func() error {
		// untyped nil
		var e error
		return e
	}()
	isUntyped, isTyped = IsNil(e3i)
	fmt.Printf("e3i == nil: %v\n", e3i == nil)                             // e3i == nil: true
	fmt.Printf("IsNil(e3i) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(e3i) -> untyped=true, typed=false

	// go sucks.
	// literal untyped nil to concrete type directly returns from a function to a interface type is typed nil.
	var e3c error = func() error {
		// untyped nil
		var e *runtime.TypeAssertionError
		return e
	}()
	isUntyped, isTyped = IsNil(e3c)
	fmt.Printf("e3c == nil: %v\n", e3c == nil)                             // e3c == nil: false
	fmt.Printf("IsNil(e3c) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(e3c) -> untyped=false, typed=true

	// literal untyped nil to concrete type directly returns from a function to a concrete type is typed nil.
	var n3 *runtime.TypeAssertionError = func() *runtime.TypeAssertionError {
		// untyped nil
		var n *runtime.TypeAssertionError
		return n
	}()
	isUntyped, isTyped = IsNil(n3)
	fmt.Printf("n3 == nil: %v\n", n3 == nil)                              // n3 == nil: true
	fmt.Printf("IsNil(n3) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(n3) -> untyped=false, typed=true

	// interface type initialized with literal untyped nil is untyped nil.
	var e4l error = nil
	isUntyped, isTyped = IsNil(e4l)
	fmt.Printf("e4l == nil: %v\n", e4l == nil)                             // e4l == nil: true
	fmt.Printf("IsNil(e4l) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(e4l) -> untyped=true, typed=false

	// interface type initialized with interface type nil is untyped nil.
	var e4i error = error(nil)
	isUntyped, isTyped = IsNil(e4i)
	fmt.Printf("e4i == nil: %v\n", e4i == nil)                             // e4i == nil: true
	fmt.Printf("IsNil(e4i) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(e4i) -> untyped=true, typed=false

	// interface type initialized with concrete type nil is typed nil.
	var e4c error = (*runtime.TypeAssertionError)(nil)
	isUntyped, isTyped = IsNil(e4c)
	fmt.Printf("e4c == nil: %v\n", e4c == nil)                             // e4c == nil: false
	fmt.Printf("IsNil(e4c) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(e4c) -> untyped=false, typed=true

	// concrete type initialized with literal untyped nil is untyped nil.
	var n4l *runtime.TypeAssertionError = nil
	isUntyped, isTyped = IsNil(n4l)
	fmt.Printf("n4l == nil: %v\n", n4l == nil)                             // n4l == nil: true
	fmt.Printf("IsNil(n4l) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(n4l) -> untyped=false, typed=true

	// concrete type initialized with concrete type nil is untyped nil.
	var n4c *runtime.TypeAssertionError = (*runtime.TypeAssertionError)(nil)
	isUntyped, isTyped = IsNil(n4c)
	fmt.Printf("n4c == nil: %v\n", n4c == nil)                             // n4c == nil: true
	fmt.Printf("IsNil(n4c) -> untyped=%v, typed=%v\n", isUntyped, isTyped) // IsNil(n4c) -> untyped=false, typed=true

	// Output:
	// e1 == nil: true
	// IsNil(e1) -> untyped=true, typed=false
	// n1 == nil: true
	// IsNil(n1) -> untyped=false, typed=true
	// e2 == nil: true
	// IsNil(e2) -> untyped=true, typed=false
	// n2 == nil: true
	// IsNil(n2) -> untyped=false, typed=true
	// e3i == nil: true
	// IsNil(e3i) -> untyped=true, typed=false
	// e3c == nil: false
	// IsNil(e3c) -> untyped=false, typed=true
	// n3 == nil: true
	// IsNil(n3) -> untyped=false, typed=true
	// e4l == nil: true
	// IsNil(e4l) -> untyped=true, typed=false
	// e4i == nil: true
	// IsNil(e4i) -> untyped=true, typed=false
	// e4c == nil: false
	// IsNil(e4c) -> untyped=false, typed=true
	// n4l == nil: true
	// IsNil(n4l) -> untyped=false, typed=true
	// n4c == nil: true
	// IsNil(n4c) -> untyped=false, typed=true
}

func ExampleIsNil_how() {
	var err error

	fmt.Printf("err == nil: %v\n", err == nil) // err == nil: true
	{
		// err is an untyped nil now.
		u, t := IsNil(err)
		fmt.Printf("IsNil(err) -> %v, %v\n", u, t) // IsNil(err) -> true, false
	}

	err = func() error {
		return nil
	}()

	fmt.Printf("err == nil: %v\n", err == nil) // err == nil: true
	{
		// err is an untyped nil now.
		u, t := IsNil(err)
		fmt.Printf("IsNil(err) -> %v, %v\n", u, t) // IsNil(err) -> true, false
	}

	err = func() error {
		var n *runtime.TypeAssertionError = nil
		return n
	}()

	fmt.Printf("err == nil: %v\n", err == nil) // err == nil: false
	{
		// err is a typed nil now.
		// It's very hard to check a typed nil is nil with a literal untyped nil.
		u, t := IsNil(err)
		fmt.Printf("IsNil(err) -> %v, %v\n", u, t) // IsNil(err) -> false, true
	}

	// Output:
	// err == nil: true
	// IsNil(err) -> true, false
	// err == nil: true
	// IsNil(err) -> true, false
	// err == nil: false
	// IsNil(err) -> false, true
}
