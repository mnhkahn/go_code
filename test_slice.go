package main

import (
	"fmt"
)

func main() {
	ori_slice := []int{1, 2, 3, 4, 5}
	fmt.Println(ori_slice)
	fmt.Println(&ori_slice[1])
	new_slice := ori_slice[1:]
	fmt.Println(&new_slice[0])
}
