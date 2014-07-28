package main

import (
	"github.com/willf/bloom"
)

func main() {
	var n uint = 1000
	filter := bloom.New(20*n, 5) // load of 20, 5 keys
	filter.Add([]byte("Love"))

	if filter.Test([]byte("Love")) {
		println("It is in black list!")
	}
}
