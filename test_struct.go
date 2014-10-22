package main

import "fmt"

type viewstruct struct {
	test int
}

type view struct {
	viewint       int
	viewstring    string
	viewbool      bool
	viewfloat32   float32
	viewarray     []int
	viewstruct    viewstruct
	viewstructarr []viewstruct
	viewpt        *viewstruct
}

func main() {
	v1 := view{}
	v2 := new(view)
	fmt.Println(v1)
	fmt.Println(v2)
}
