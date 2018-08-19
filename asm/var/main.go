// Package _var
package _var

var Id = 1031

/*
go tool compile -S pkg.go
"".Id SNOPTRDATA size=8
	0x0000 07 04 00 00 00 00 00 00                          ........
*/

// size int 是 8 个字节
// 十六进制 0x407
