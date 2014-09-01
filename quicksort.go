package main

import "fmt"

type BySortIndex []int

func (a BySortIndex) Len() int      { return len(a) }
func (a BySortIndex) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySortIndex) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	test := []int{49, 38, 65, 97, 76, 13, 27, 49}
	quickSort(BySortIndex(test), 0, len(test))
	fmt.Println(test)
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// abc分别是数组头、中间和尾三个元素脚标
// 将中间元素放在头，最小的元素放在中间，最大的元素放在最后
func medianOfThree(data Interface, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	if data.Less(m2, m1) {
		data.Swap(m2, m1)
	}
	if data.Less(m1, m0) {
		data.Swap(m1, m0)
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func swapRange(data Interface, a, b, n int) {
	for i := 0; i < n; i++ {
		data.Swap(a+i, b+i)
	}
	// fmt.Println(data, "---------------")
}

func doPivot(data Interface, lo, hi int) (midlo, midhi int) {
	// m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	// if hi-lo > 40 {
	// 	// Tukey's ``Ninther,'' median of three medians of three.
	// 	s := (hi - lo) / 8
	// 	medianOfThree(data, lo, lo+s, lo+2*s)
	// 	medianOfThree(data, m, m-s, m+s)
	// 	medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	// }
	// medianOfThree(data, lo, m, hi-1)

	// 根据前面的准备，数组头存放的是中间元素，作为快速排序的轴
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		// 遍历，找到第一个大于轴的元素
		for b < c {
			if data.Less(b, pivot) { // data[b] < pivot
				b++
			} else if !data.Less(pivot, b) { // data[b] = pivot
				data.Swap(a, b)
				a++
				b++
			} else {
				break
			}
		}
		// 遍历，找到第一个小于轴的元素
		for b < c {
			if data.Less(pivot, c-1) { // data[c-1] > pivot
				c--
			} else if !data.Less(c-1, pivot) { // data[c-1] = pivot
				data.Swap(c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] < pivot
		data.Swap(b, c-1)
		b++
		c--
		// fmt.Println(data, "**********")
	}
	n := min(b-a, a-lo)
	swapRange(data, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRange(data, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSort(data Interface, a, b int) {
	if b-a <= 1 {
		return
	}

	mlo, mhi := doPivot(data, a, b)
	quickSort(data, a, mlo)
	quickSort(data, mhi, b)
}
