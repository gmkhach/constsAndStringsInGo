package main

import (
	"fmt"
	"runtime"
)

// Here are some untyped constants. These are called constants of a kind.
// The compiler has more flexibility with untyped constants.
const a = 42
const b = 3.14
const c = "Am I not merciful!?"

// Now let's declare some typed constants.
// We will also use a multiple constant declaration this time just to demonstrate.
const (
	d int     = 64
	e float64 = 2.7
	f string  = "... and I will have my vengence, in this life or the next"
)

func main() {

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	// Strings are immutable in GO
	// Strings are actually slices of bytes
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// Printing the UTF8 values for each byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
