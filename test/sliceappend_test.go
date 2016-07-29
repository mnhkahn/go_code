package test

import "testing"

// go test -bench=Slice* -benchmem

func BenchmarkSliceAppend(b *testing.B) {
	a := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		a = append(a, i)
	}
}

func BenchmarkSliceSet(b *testing.B) {
	a := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		a[i] = i
	}
}

/*
BenchmarkSliceAppend-4  200000000                7.87 ns/op            8 B/op          0 allocs/op
BenchmarkSliceSet-4     300000000                5.76 ns/op            8 B/op          0 allocs/op
*/
