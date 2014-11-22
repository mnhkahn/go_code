package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Split("", ";")
	fmt.Printf("%d****%s****\n", len(a), a[0])
}
