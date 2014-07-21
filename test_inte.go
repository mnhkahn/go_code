package main

type Phone interface {
	Company() string
}

type Android struct {
}

func (p *Android) Company() string {
	return "Google"
}

type iPhone struct {
}

func (p *iPhone) Company() string {
	return "Apple"
}

func main() {
	android := Android{}
	iphone := iPhone{}
	println(android.Company())
	println(iphone.Company())
}
