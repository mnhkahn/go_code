package main

func a() int { return 1 }

func main() {
	print(a())
}

/*
go build -gcflags="-N -l"
go tool objdump -s "main\.a" mock1
TEXT main.a(SB) /Users/lichao/code/gocode/src/github.com/mnhkahn/go_code/mock/mock1/main.go
  main.go:3		0x104aad0		48c744240800000000	MOVQ $0x0, 0x8(SP)
  main.go:3		0x104aad9		48c744240801000000	MOVQ $0x1, 0x8(SP)
  main.go:3		0x104aae2		c3			RET
  :-1			0x104aae3		cc			INT $0x3
  :-1			0x104aae4		cc			INT $0x3
  :-1			0x104aae5		cc			INT $0x3
  :-1			0x104aae6		cc			INT $0x3
  :-1			0x104aae7		cc			INT $0x3
  :-1			0x104aae8		cc			INT $0x3
  :-1			0x104aae9		cc			INT $0x3
  :-1			0x104aaea		cc			INT $0x3
  :-1			0x104aaeb		cc			INT $0x3
  :-1			0x104aaec		cc			INT $0x3
  :-1			0x104aaed		cc			INT $0x3
  :-1			0x104aaee		cc			INT $0x3
  :-1			0x104aaef		cc			INT $0x3
go tool objdump -s "main\.main"  mock1
TEXT main.main(SB) /Users/lichao/code/gocode/src/github.com/mnhkahn/go_code/mock/mock1/main.go
  main.go:5		0x104aaf0		65488b0c25a0080000	MOVQ GS:0x8a0, CX
  main.go:5		0x104aaf9		483b6110		CMPQ 0x10(CX), SP
  main.go:5		0x104aafd		763e			JBE 0x104ab3d
  main.go:5		0x104aaff		4883ec18		SUBQ $0x18, SP
  main.go:5		0x104ab03		48896c2410		MOVQ BP, 0x10(SP)
  main.go:5		0x104ab08		488d6c2410		LEAQ 0x10(SP), BP
  main.go:6		0x104ab0d		e8beffffff		CALL main.a(SB)
  main.go:6		0x104ab12		488b0424		MOVQ 0(SP), AX
  main.go:6		0x104ab16		4889442408		MOVQ AX, 0x8(SP)
  main.go:6		0x104ab1b		e8b06afdff		CALL runtime.printlock(SB)
  main.go:6		0x104ab20		488b442408		MOVQ 0x8(SP), AX
  main.go:6		0x104ab25		48890424		MOVQ AX, 0(SP)
  main.go:6		0x104ab29		e82272fdff		CALL runtime.printint(SB)
  main.go:6		0x104ab2e		e81d6bfdff		CALL runtime.printunlock(SB)
  main.go:7		0x104ab33		488b6c2410		MOVQ 0x10(SP), BP
  main.go:7		0x104ab38		4883c418		ADDQ $0x18, SP
  main.go:7		0x104ab3c		c3			RET
  main.go:5		0x104ab3d		e8de88ffff		CALL runtime.morestack_noctxt(SB)
  main.go:5		0x104ab42		ebac			JMP main.main(SB)
  :-1			0x104ab44		cc			INT $0x3
  :-1			0x104ab45		cc			INT $0x3
  :-1			0x104ab46		cc			INT $0x3
  :-1			0x104ab47		cc			INT $0x3
  :-1			0x104ab48		cc			INT $0x3
  :-1			0x104ab49		cc			INT $0x3
  :-1			0x104ab4a		cc			INT $0x3
  :-1			0x104ab4b		cc			INT $0x3
  :-1			0x104ab4c		cc			INT $0x3
  :-1			0x104ab4d		cc			INT $0x3
  :-1			0x104ab4e		cc			INT $0x3
  :-1			0x104ab4f		cc			INT $0x3

真SP寄存器对应的是栈的顶部
*/
