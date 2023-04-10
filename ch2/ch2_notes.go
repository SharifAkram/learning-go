package main

import "fmt"

// Chapter 2 - Primitive Types and Declarations

// boolean

var flag bool // no value assigned, set to false
var isAwesome = true

// numeric types

int8
int16
int32
int64
uint8
unit16
unit32
unit64

// var Versus :=

var x int = 10
var x = 10
var x int // zero value
var x, y int = 10, 20
var x, y int           // zero values of the same type
var x, y = 10, "hello" // different types

// declaring multiple variables at once, wrap them in a declaration list:


var (
	x    int
	y        = 20
	z    int = 30
	d, e     = 40, "hello"
	f, g string
)

//var x = 10
x := 10

//var x, y = 10, "hello"
x, y := 10, "hello"

// assign values to existing variables
x := 10
x, y := 30, "hello"

// There is one limitation on :=. If you are declaring a variable at
// package level, you must use var because := is not legal outside of
// functions.

// untyped constant declaration:
const x = 10

// legal assignments:
var y int = x
var z float64 = x
var d byte = x

// typed declaration:
const typedX int = 10

// unused variables:

func main() {
	x := 10
	x = 20
	fmt.Println(x)
	x = 30
}
