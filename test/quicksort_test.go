package test

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	test0 := []int{49, 38, 65, 97, 76, 13, 27, 49}
	test1 := []int{13, 27, 38, 49, 49, 65, 76, 97}
	heapSort(BySortIndex(test0), 0, len(test0))
	for i := 0; i < len(test0); i++ {
		if test0[i] != test1[i] {
			t.Fatal("error")
		}
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.StopTimer()  //调用该函数停止压力测试的时间计数
	b.StartTimer() //重新开始时间

	b.N = 1234

	for i := 0; i < b.N; i++ {
		test0 := []int{49, 38, 65, 97, 76, 13, 27, 49}
		test1 := []int{13, 27, 38, 49, 49, 65, 76, 97}
		heapSort(BySortIndex(test0), 0, len(test0))
		for i := 0; i < len(test0); i++ {
			if test0[i] != test1[i] {
				b.Fatal("error")
			}
		}
	}
}
