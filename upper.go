package main

import "fmt"

const MaxASCII = '\u007F'

func toUpper(r rune) rune {
	if r <= MaxASCII {
		if 'a' <= r && r <= 'z' {
			r -= 'a' - 'A'
		}
		return r
	}
	return r
}

func ToUpper(s []rune) (res []rune) {
	for i := 0; i < len(s); i++ {
		res = append(res, toUpper(s[i]))
	}
	return res
}

func main() {
	a := "Hello, 世界"
	fmt.Println(string(ToUpper([]rune(a))))

}
