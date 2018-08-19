// Package funcaddr
package main

import (
	"fmt"
	"unsafe"
)

func a() int { return 1 }

func main() {
	f := a
	//fmt.Printf("0x%x\n", *(*uintptr)(unsafe.Pointer(&f)))
	fmt.Printf("0x%x\n", **(**uintptr)(unsafe.Pointer(&f)))
}

/*
go build -gcflags="-N -l"
go tool objdump -s "main\.a"  funcaddr
TEXT main.a(SB) /Users/lichao/code/gocode/src/github.com/mnhkahn/go_code/mock/funcaddr/main.go
  main.go:9		0x108eed0		48c744240800000000	MOVQ $0x0, 0x8(SP)
  main.go:9		0x108eed9		48c744240801000000	MOVQ $0x1, 0x8(SP)
  main.go:9		0x108eee2		c3			RET
  :-1			0x108eee3		cc			INT $0x3
  :-1			0x108eee4		cc			INT $0x3
  :-1			0x108eee5		cc			INT $0x3
  :-1			0x108eee6		cc			INT $0x3
  :-1			0x108eee7		cc			INT $0x3
  :-1			0x108eee8		cc			INT $0x3
  :-1			0x108eee9		cc			INT $0x3
  :-1			0x108eeea		cc			INT $0x3
  :-1			0x108eeeb		cc			INT $0x3
  :-1			0x108eeec		cc			INT $0x3
  :-1			0x108eeed		cc			INT $0x3
  :-1			0x108eeee		cc			INT $0x3
  :-1			0x108eeef		cc			INT $0x3
./funcaddr
0x10c6cb0
*/

/*
go build -gcflags="-N -l"
./funcaddr
0x108eed0
go tool objdump -s "main\.a"  funcaddr
TEXT main.a(SB) /Users/lichao/code/gocode/src/github.com/mnhkahn/go_code/mock/funcaddr/main.go
  main.go:9		0x108eed0		48c744240800000000	MOVQ $0x0, 0x8(SP)
  main.go:9		0x108eed9		48c744240801000000	MOVQ $0x1, 0x8(SP)
  main.go:9		0x108eee2		c3			RET
  :-1			0x108eee3		cc			INT $0x3
  :-1			0x108eee4		cc			INT $0x3
  :-1			0x108eee5		cc			INT $0x3
  :-1			0x108eee6		cc			INT $0x3
  :-1			0x108eee7		cc			INT $0x3
  :-1			0x108eee8		cc			INT $0x3
  :-1			0x108eee9		cc			INT $0x3
  :-1			0x108eeea		cc			INT $0x3
  :-1			0x108eeeb		cc			INT $0x3
  :-1			0x108eeec		cc			INT $0x3
  :-1			0x108eeed		cc			INT $0x3
  :-1			0x108eeee		cc			INT $0x3
  :-1			0x108eeef		cc			INT $0x3
*/
