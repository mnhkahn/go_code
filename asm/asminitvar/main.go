// Package asminitvar
package main

import "../pkg"

func main() {
	println(pkg.Id)

	a, b := pkg.Swap(1, 2)
	println(a, b)
}
