package main

import "log"

func Read(buf []byte) (int, error) {
    b:=[]byte{} 
	b = append(b, byte('a'))
n:=copy(buf,b)
log.Println(buf,"AAA")
    return n,nil
}

func main() {
	var a = make([]byte,10)
	log.Println(a)
	Read(a)
	log.Println(a)
}
