package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7}
	println(&a)
	println(&b)

	c := make([]int, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	fmt.Println(c)
	println(&c)

	d := append(a, b...)
	fmt.Println(d)
	println(&d)
}
