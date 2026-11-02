package main

import (
	"fmt"
	"strconv"
)

var f float32 = 36

func main() {
	fmt.Println("Hello, my first go program")
	fmt.Println("Variable declaration:")
	var i int
	i = 43
	// println(f)// this will print the package level declaration
	// var f float32 = 33 // if we uncomment this, f value will be this as innermost scope takes precedence in declaration this is called shadowing
	m := 25
	fmt.Println(i, f, m)
	//fmt.Printf("%v, %T", f, f)
	fmt.Printf("%v, %T\n", m, m)
	var n float32 = float32(m)
	fmt.Printf("%v, %T\n", n, n)

	var j string
	j = strconv.Itoa(m)
	fmt.Printf("%v, %T\n", j, j)
}
