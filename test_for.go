package main

import "fmt"

func main() {
	a := []string{}
	a = append(a, "hello")
	a = append(a, ", ")
	a = append(a, "world")
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(&a[i])
	}

	for _, aa := range a {
		aa += "@"
		fmt.Println(&aa)
	}
	for i := 0; i < len(a); i++ {
		a[i] += "@"
		fmt.Println(&a[i])
	}
	fmt.Println(a)
}
