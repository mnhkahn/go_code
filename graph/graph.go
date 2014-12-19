package graph

import (
	"fmt"
)

type Graph struct {
	edgnum int
	vexnum int
	adj    []Vertex
}

type Vertex struct {
	Data string
	e    *Edge
}

type Edge struct {
	ivex int
	next *Edge
}

func Create() *Graph {
	g := new(Graph)
	g.adj = make([]Vertex, 0)
	return g
}

func NewVertex(data string) *Vertex {
	return &Vertex{Data: data}
}

func NewEdge(i int) *Edge {
	return &Edge{ivex: i}
}

func (this *Graph) InsertVertex(v Vertex) {
	this.adj = append(this.adj, v)
}

func (this *Graph) InsertEdge(v1, v2 Vertex) {
	var p *Edge = this.adj[this.get_position(v1.Data)].e
	if p == nil {
		// fmt.Println("nil...", v1.Data, v2.Data)
		this.adj[this.get_position(v1.Data)].e = NewEdge(this.get_position(v2.Data))
	} else {
		for ; p.next != nil; p = p.next {
		}
		p.next = NewEdge(this.get_position(v2.Data))
		// fmt.Println("append...", v1.Data, v2.Data)
	}
}

func (this *Graph) DeleteVertex(v Vertex) {

}

func (this *Graph) DeleteEdge(v1, v2 Vertex) {

}

func (this *Graph) IsEmpty() bool {
	if len(this.adj) == 0 {
		return true
	}
	return false
}

func (this *Graph) Adjacent() []Vertex {
	return this.adj
}

func (this *Graph) Traverse() []Vertex {
	return nil
}

func (this *Graph) dfs() []Vertex {
	return nil
}

func (this *Graph) bfs() []Vertex {
	return nil
}

/*
 * 打印邻接表图
 */
func (this *Graph) Print() {
	var node *Edge = nil

	fmt.Printf("List Graph:\n")
	for i := 0; i < this.vexnum; i++ {
		fmt.Printf("%d(%s): ", i, this.adj[i].Data)
		node = this.adj[i].e
		// fmt.Println("*************", node)
		for node != nil {
			fmt.Printf("%d(%s) ", node.ivex, this.adj[node.ivex].Data)
			node = node.next
		}
		fmt.Printf("\n")
	}
}

/*
 * 返回ch在matrix矩阵中的位置
 */
func (this *Graph) get_position(data string) int {
	i := -1
	for i = 0; i < this.vexnum; i++ {
		if this.adj[i].Data == data {
			return i
		}
	}
	return -1
}
