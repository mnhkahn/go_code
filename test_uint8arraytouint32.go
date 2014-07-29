package main

import "fmt"
import "encoding/binary"

func main() {
	var a []byte = []byte{0, 1, 2, 3}
	fmt.Println(a)
	fmt.Println(binary.BigEndian.Uint32(a))
	fmt.Println(binary.LittleEndian.Uint32(a))
}
