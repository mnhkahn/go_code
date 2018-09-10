package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

var _ fmt.Formatter = &data{}

type data struct {
	A string
	B int
}

func (d *data) Format(f fmt.State, c rune) {
	switch c {
	case 'v': // &{1 2}
		buf, err := json.Marshal(d)
		if err != nil {
			panic(err)
		}
		f.Write(buf)
	case 's':
		f.Write([]byte(d.String()))
	case 'x', 'X':
		//case 'p':
		v := reflect.ValueOf(d)
		f.Write([]byte{'('})
		f.Write([]byte(v.Type().String()))
		f.Write([]byte{')', '('})
		u := v.Pointer()
		f.Write([]byte(strconv.FormatUint(uint64(u), 16)))
		f.Write([]byte{')'})
	default:
		f.Write([]byte("http://cyeam.com"))
	}
}

func (d *data) String() string {
	buf, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func main() {
	d := &data{"1", 2}
	fmt.Printf("v %v\n", d)
	fmt.Printf("s %s\n", d)
	fmt.Printf("p %p\n", d)
	fmt.Printf("T %T\n", d)
	fmt.Printf("b %b\n", d)
	fmt.Printf("o %o\n", d)
	fmt.Printf("x %x\n", d)
	fmt.Printf("d %d\n", d)

}

/*
宽度，精度
p 和 T 是自动实现，无法复写
	switch verb {
	case 'T':
		p.fmt.fmtS(reflect.TypeOf(arg).String())
		return
	case 'p':
		p.fmtPointer(reflect.ValueOf(arg), 'p')
		return
	}

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p

%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0

	case '-':
		return p.fmt.minus
	case '+':
		return p.fmt.plus || p.fmt.plusV
	case '#':
		return p.fmt.sharp || p.fmt.sharpV
	case ' ':
		return p.fmt.space
	case '0':
		return p.fmt.zero
	}

+	always print a sign for numeric values;
	guarantee ASCII-only output for %q (%+q)
-	pad with spaces on the right rather than the left (left-justify the field)
#	alternate format: add leading 0 for octal (%#o), 0x for hex (%#x);
	0X for hex (%#X); suppress 0x for %p (%#p);
	for %q, print a raw (backquoted) string if strconv.CanBackquote
	returns true;
	always print a decimal point for %e, %E, %f, %F, %g and %G;
	do not remove trailing zeros for %g and %G;
	write e.g. U+0078 'x' if the character is printable for %U (%#U).
' '	(space) leave a space for elided sign in numbers (% d);
	put spaces between bytes printing strings or slices in hex (% x, % X)
0	pad with leading zeros rather than spaces;
	for numbers, this moves the padding after the sign

defer p.catchPanic(p.arg, verb)
 */
