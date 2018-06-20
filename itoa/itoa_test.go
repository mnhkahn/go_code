package itoa

import (
	"testing"
	"strconv"
	"fmt"
)

var smallInt = 35
var bigInt = 999999999999999

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(smallInt)
		_ = val
	}
}

func BenchmarkItoaFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatInt(int64(smallInt), 10)
		_ = val
	}
}

func BenchmarkItoaSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%d", smallInt)
		_ = val
	}
}

func BenchmarkItoaBase2Sprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%b", smallInt)
		_ = val
	}
}

func BenchmarkItoaBase2FormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatInt(int64(smallInt), 2)
		_ = val
	}
}

func BenchmarkItoaBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(bigInt)
		_ = val
	}
}

func BenchmarkItoaFormatIntBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := strconv.FormatInt(int64(bigInt), 10)
		_ = val
	}
}

func BenchmarkItoaSprintfBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := fmt.Sprintf("%d", bigInt)
		_ = val
	}
}
