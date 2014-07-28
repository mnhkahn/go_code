package main

import (
	"crypto/md5"
	"encoding/hex"
)

func main() {
	m := md5.New()
	m.Write([]byte("hello, world"))
	s := hex.EncodeToString(m.Sum(nil))
	println(s)
}
