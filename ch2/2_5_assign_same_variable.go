// look alike variable

package main

import "fmt"

func main() {
	ａ := "hello"   // Unicode U+FF41
	a := "goodbye" // standard lowercase a (Unicode U+0061),
	// different from the line above!
	fmt.Println(ａ)
	fmt.Println(a)
}
