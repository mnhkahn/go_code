GLOBL ·Id(SB),$8

DATA ·Id+0(SB)/1,$0x07
DATA ·Id+1(SB)/1,$0x04
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00

TEXT ·Swap(SB), $0-32
    MOVQ a+0(FP),     AX // a
    MOVQ b+8(FP),     BX // b
    MOVQ ret0+16(FP), CX // ret0
    MOVQ ret1+24(FP), DX // ret1
    MOVQ AX, ret1+24(FP)
    MOVQ BX, ret0+16(FP)
    RET

