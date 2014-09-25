package test

import "fmt"
import "sort"

type BySortIndex []int

func (a BySortIndex) Len() int      { return len(a) }
func (a BySortIndex) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySortIndex) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	test0 := []int{49, 38, 65, 97, 76, 13, 27, 49}
	test1 := []int{49, 38, 65, 97, 76, 13, 27, 49}
	fmt.Println(test0)
	sort.Sort(BySortIndex(test0))
	fmt.Println(test0)

	heapSort(BySortIndex(test1), 0, len(test1))
	fmt.Println(test1)

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

func heapSort(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	// 二叉树结构当中最后一个有子结点的结点
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}

// 建立树函数
// 父节点root的左孩子2*root + 1
func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi { // child 超出了数组长度，也就是该结点无孩子结点，返回
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) { // 右孩子结点存在
			child++
		}
		// 以上是为了在孩子结点当中找到较大的结点，用child表示
		if !data.Less(first+root, first+child) {
			return
		}
		// 如果根结点小于较大的孩子结点，则进行交换；该孩子结点的堆结构可能会被影响，继续去处理孩子结点
		data.Swap(first+root, first+child)
		root = child
	}
}
