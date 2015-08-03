/*
https://godoc.org/github.com/yuin/gopher-lua
/*

package main

import (
	"github.com/yuin/gopher-lua"

	"go_code/lua/mymodule"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("mymodule", mymodule.Loader)
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
}
