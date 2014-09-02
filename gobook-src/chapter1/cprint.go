package main

/*
#include <stdio.h>
*/
import "C"

func main() {
	cstr := C.CString("Hello, world")
    C.puts(cstr)
}
