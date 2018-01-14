package main

import "fmt"

func prec() {
	x := 2
	y := 3
	z := x << uint(4+y) << 2
	fmt.Printf("%s %d\n", "hello", z)
}
