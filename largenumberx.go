package main

import (
	"fmt"
	"github.com/mnhkahn/go-util/strings"
)

func LargeNumberMultiplication(a string, b string) (reuslt string) {
	a = strings.Reverse(a)
	b = strings.Reverse(b)
	c := make([]byte, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i+j] += (a[i] - '0') * (b[j] - '0')
		}
	}

	var plus byte = 0
	for i := 0; i < len(c); i++ {
		if c[i] == 0 {
			break
		}
		temp := c[i] + plus
		plus = 0
		if temp > 9 {
			plus = temp / 10
			reuslt += string(temp - plus*10 + '0')
		} else {
			reuslt += string(temp + '0')
		}

	}
	return strings.Reverse(reuslt)
}

func main() {
	a := "12345633333333333333333"
	// b := "123"
	// fmt.Println(LargeNumberMultiplication(a, b))
	// fmt.Println(LargeNumberMultiplication(a, "1"))
	fmt.Println(LargeNumberMultiplication(a, "123"))
	// fmt.Println(LargeNumberMultiplication(a, "3"))
}
