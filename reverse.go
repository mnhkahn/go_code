package main

func Reverse(s string) string {
	r := []rune(s)
	return string(r)
}
func main() {
	a := "Hello, 世界"
	println(Reverse(a))
}
