package main

import (
	"fmt"

	"github.com/willf/bitset"
)

func main() {
	var b bitset.BitSet // 定义一个BitSet对象
	fmt.Println(b.Bytes())
    b.Set(0)
    fmt.Println(b.Bytes(),0)
	b.Set(10) // 给这个set新增两个值10
    fmt.Println(b.Bytes(),0,10)
    b.Set(64)
	fmt.Println(b.Bytes(),0,10,64)
	if b.Test(1000) { // 查看set中是否有1000这个值（我觉得Test这个名字起得是真差劲，为啥不叫Exist）
		b.Clear(1000) // 情况set
	}
}
