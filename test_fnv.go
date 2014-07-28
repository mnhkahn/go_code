package main

import "fmt"
import "hash/fnv"
import "encoding/hex"

func main() {
	a := fnv.New32()
	a.Write([]byte("hello"))
	fmt.Println(hex.EncodeToString(a.Sum(nil)))
}
