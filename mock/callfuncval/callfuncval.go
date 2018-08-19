// Package callfuncval
package main

func a() int { return 1 }

func main() {
	f := a
	f()
}

/*
go build -gcflags="-N -l"
go tool objdump -s "main\.main" callfuncval
TEXT main.main(SB) /Users/lichao/code/gocode/src/github.com/mnhkahn/go_code/mock/callfuncval/callfuncval.go
  callfuncval.go:6	0x104aaf0		65488b0c25a0080000	MOVQ GS:0x8a0, CX
  callfuncval.go:6	0x104aaf9		483b6110		CMPQ 0x10(CX), SP
  callfuncval.go:6	0x104aafd		762d			JBE 0x104ab2c
  callfuncval.go:6	0x104aaff		4883ec18		SUBQ $0x18, SP
  callfuncval.go:6	0x104ab03		48896c2410		MOVQ BP, 0x10(SP)
  callfuncval.go:6	0x104ab08		488d6c2410		LEAQ 0x10(SP), BP
  callfuncval.go:7	0x104ab0d		488d1554010200		LEAQ go.func.*+58(SB), DX
  callfuncval.go:7	0x104ab14		4889542408		MOVQ DX, 0x8(SP)
  callfuncval.go:8	0x104ab19		488b0548010200		MOVQ go.func.*+58(SB), AX
  callfuncval.go:8	0x104ab20		ffd0			CALL AX
  callfuncval.go:9	0x104ab22		488b6c2410		MOVQ 0x10(SP), BP
  callfuncval.go:9	0x104ab27		4883c418		ADDQ $0x18, SP
  callfuncval.go:9	0x104ab2b		c3			RET
  callfuncval.go:6	0x104ab2c		e8ef88ffff		CALL runtime.morestack_noctxt(SB)
  callfuncval.go:6	0x104ab31		ebbd			JMP main.main(SB)
  :-1			0x104ab33		cc			INT $0x3
  :-1			0x104ab34		cc			INT $0x3
  :-1			0x104ab35		cc			INT $0x3
  :-1			0x104ab36		cc			INT $0x3
  :-1			0x104ab37		cc			INT $0x3
  :-1			0x104ab38		cc			INT $0x3
  :-1			0x104ab39		cc			INT $0x3
  :-1			0x104ab3a		cc			INT $0x3
  :-1			0x104ab3b		cc			INT $0x3
  :-1			0x104ab3c		cc			INT $0x3
  :-1			0x104ab3d		cc			INT $0x3
  :-1			0x104ab3e		cc			INT $0x3
  :-1			0x104ab3f		cc			INT $0x3
*/
