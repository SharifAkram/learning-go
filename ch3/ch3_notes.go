package main

// arrays - too rigid to use directly

// (x[0], x[1], and x[2])
var x [3]int

// array literal
var x = [3]int{10, 20, 30}
var x = [...]int{10, 20, 30}

// sparse array
// [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
var x = [12]int{1, 5: 4, 6, 10: 100, 15}

// == and !=
var x = [...]int{1, 2, 3}
var y = [3]int{1, 2, 3}
fmt.Println(x == y)				// prints true

// multi-dimensional array
var x [2][3]int

x[0] = 10
fmt.Println(x[2])
fmt.Println(len(x))

// Slices - similar to arrays in other languages
var x = []int{10, 29, 30}

// [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15].
var x = []int{1, 5: 4, 6, 10: 100, 15}

// multidimensional slices
var x [][]int

// nil slice

var x []int
fmt.Println(x == nil) // prints true

// append
var x []int
x = append(x, 10)

var x = []int{1, 2, 3}
x = append(x, 4)

// more than one value at a time
x = append(x, 5, 6, 7)

// y := [](20, 30, 40)
// x = append(x, y...)

// make
x := make([]int, 5)

// DOES NOT WORK
x := make([]int, 5)
x = append(x, 10)

// specify initial capacity
x := make([]int, 5, 10)

// create a slice with zero length, but capacity greater than 0
// non nil slice with a length of zero, capacity of 10
// cannot index into it, but append
x := make([]int, 0, 10)
x = append(x, 5, 6, 7, 8)

// The value of x is now [5 6 7 8], with a length of 4 and a capacity of 10.

// declaring a slice that might stay nil
var data []int

// declaring a slice with default values
data := []int{2, 4, 6, 8}
