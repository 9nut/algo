package AVL

import (
	"fmt"
	"github.com/9nut/algo/pq"
	"testing"
)

type IntType int

func (I IntType) Equal(j Comparable) bool {
	v, ok := j.(IntType)
	if !ok {
		panic("Equal called with wrong type")
	}
	return I == v
}

func (I IntType) Less(j Comparable) bool {
	v, ok := j.(IntType)
	if !ok {
		panic("Less called with wrong type")
	}
	return I < v
}

// breadth first search
func BFS(b *Node, pq *PQ.PQ, level int) {
	if b != nil {
		pq.Push(level, *b)
		BFS(b.l, pq, level+1)
		BFS(b.r, pq, level+1)
	}
}

func TestAVL(*testing.T) {
	tree := NewNode(IntType(0))
	fmt.Println("Inserting nodes")
	for i := 0; i <= 15; i++ {
		fmt.Println("Inserting ", i)
		tree = tree.Insert(IntType(i))
	}
	prnode := func(v interface{}) {
		vv, ok := v.(IntType)
		if ok {
			fmt.Println(vv)
		}
	}
	fmt.Println("In-order traversal...")
	tree.Traverse(prnode)
	fmt.Println("Rerverse-order traversal...")
	tree.RTraverse(prnode)

	fmt.Println("Search...")
	for i := 0; i <= 15; i++ {
		c, p := tree.Find(IntType(i))
		fmt.Println("Find(", i, "): c{", c, "}", " p{", p, "}")
	}

	// print breadths first
	var q PQ.PQ
	level := 0
	BFS(tree, &q, level)
	fmt.Println("Breadth first traversal...")
	for len(q) > 0 {
		n := q.Pop()
		if level != n.P {
			level = n.P
			fmt.Println("\nlevel:", level)
			for i := n.P; i >= 0; i-- {
				fmt.Print(" ")
			}
		}
		fmt.Print(n.V, " ")
	}

	fmt.Println("Removing largest...")
	for tree != nil {
		x := tree.Largest()
		fmt.Println("removing ", x)
		tree = tree.Remove(x)
	}
}

func BenchmarkAVLInsert(b *testing.B) {
	tree := NewNode(IntType(0))
	for i := 0; i < b.N; i++ {
		tree = tree.Insert(IntType(i))
	}
}

func BenchmarkAVLRemove(b *testing.B) {
	tree := NewNode(IntType(0))
	for i := 0; i < b.N; i++ {
		tree = tree.Insert(IntType(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree = tree.Remove(IntType(i))
	}
}
