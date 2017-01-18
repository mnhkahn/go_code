package main

import (
	"testing"

	"github.com/RoaringBitmap/roaring"
	"github.com/willf/bitset"
)

func BenchmarkMap(b *testing.B) {
	var B = make(map[int]int8, 3)
	B[10] = 1
	B[11] = 1
	for i := 0; i < b.N; i++ {
		if _, exists := B[1]; exists {

		}
		if _, exists := B[11]; exists {

		}
		if _, exists := B[1000000]; exists {

		}
	}
}

func BenchmarkBitset(b *testing.B) {
	var B bitset.BitSet
	B.Set(10).Set(11)
	for i := 0; i < b.N; i++ {
		if B.Test(1) {

		}
		if B.Test(11) {

		}
		if B.Test(1000000) {

		}
	}
}

func BenchmarkRoaring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B := roaring.BitmapOf(10, 11)
		if B.ContainsInt(1) {

		}
		if B.ContainsInt(11) {
		}
		if B.ContainsInt(1000000) {

		}

	}
}
