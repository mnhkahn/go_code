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

func (this *Graph) Adjacent(v Vertex) []Vertex {
	res := make([]Vertex, 0)
	p := this.adj[this.get_position(v.Data)].e
	for ; p != nil; p = p.next {
		res = append(res, this.adj[p.ivex])
	}
	return res
}

func (this *Graph) Dfs() []Vertex {
	return nil
}

func (this *Graph) Bfs() []Vertex {
	res := make([]Vertex, 0)
	for _, v := range this.adj {
		adjs := this.Adjacent(v)
		for _, a := range adjs {

		}
	}
	return res
}

/*
 * 打印邻接表图
 */
func (this *Graph) Print() {
	fmt.Printf("List Graph:\n")
	for i := 0; i < this.vexnum; i++ {
		fmt.Printf("%d(%s): ", i, this.adj[i].Data)
		adjs := this.Adjacent(this.adj[i])
		for _, a := range adjs {
			fmt.Printf("%d(%s) ", this.get_position(a.Data), a.Data)
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
