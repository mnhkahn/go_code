package main

import "fmt"

func main() {

	// 先定义一个数组备用
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 基于myArray数组创建slice，初始内容为myArray的前5个元素
	var mySlice []int = myArray[:5]
	

	fmt.Println("Elements of myArray: ")

	for v := range myArray {
		fmt.Print(v, " ")
	}
	
	
	fmt.Println("\nElements of mySlice: ")
	
	for i, v := range mySlice {
		fmt.Print(i, v, " ")
	}
	
	fmt.Println()
}
