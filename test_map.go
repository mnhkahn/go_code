package main

import (
	"fmt"
)

func main() {
	query := map[string]string{}
	// 需要按照字典排序
	query["test0"] = "0"
	query["test1"] = "1"
	query["test2"] = "2"

	for i := 0; i < 100; i++ {
		for _, v := range query {
			fmt.Print(v)
		}
		fmt.Println()
	}
}
