package main

import "fmt"

func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

type BySortIndex []int

func (a BySortIndex) Len() int      { return len(a) }
func (a BySortIndex) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySortIndex) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	test0 := []int{49, 38, 65, 97, 76, 13, 27, 49}
	insertionSort(BySortIndex(test0), 0, len(test0))
	fmt.Println(test0)
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
