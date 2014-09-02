package main

import "fmt"

func main() {
	
	a := []int{1, 2, 3, 4, 5, 6}
	
	fmt.Println("Elements of array:")	

	for v := range a {
		fmt.Print(v, " ")
	}
	
	fmt.Println()

	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		fmt.Println("i=", i, "j=", j)
		fmt.Println("a[i]=", a[i], "a[j]=", a[j])
		a[i], a[j] = a[j], a[i]
		fmt.Println("exchanged: a[i]=", a[i], "a[j]=", a[j])
	}

	fmt.Println("Elements of array:")	
	
	for _, v := range a {
		fmt.Print(v, " ")
	}
	
	fmt.Println()
}
