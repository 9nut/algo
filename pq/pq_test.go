package PQ

import (
	"fmt"
	"testing"
)

func TestPq(*testing.T) {
	J := []Node{
		{P: 1, V: 1},
		{P: 2, V: 2},
		{P: 3, V: 3},
		{P: 4, V: 4},
		{P: 5, V: 5},
		{P: 6, V: 6},
		{P: 6, V: 66},
		{P: 6, V: 666},
		{P: 7, V: 7},
		{P: 8, V: 8},
		{P: 8, V: 88},
		{P: 9, V: 9},
	}
	K := []Node{
		{P: 1, V: 1},
		{P: 2, V: 2},
		{P: 3, V: 3},
		{P: 4, V: 4},
		{P: 5, V: 5},
		{P: 6, V: 6},
		{P: 6, V: 66},
		{P: 6, V: 666},
		{P: 7, V: 7},
		{P: 8, V: 8},
		{P: 8, V: 88},
		{P: 9, V: 9},
	}
	I := PQ(J)
	I.MakePQ()
	for len(I) > 0 {
		fmt.Println(I.Pop())
	}
	for _, v := range K {
		I.Push(v.P, v)
	}
	for len(I) > 0 {
		fmt.Println(I.Pop())
	}
}

func BenchmarkPQPush(b *testing.B) {
	I := PQ{}
	for i := 0; i < b.N; i++ {
		I.Push(i, i)
	}
}

func BenchmarkPQPop(b *testing.B) {
	I := PQ{}
	for i := 0; i < b.N; i++ {
		I.Push(i, i)
	}

	b.ResetTimer()
	for len(I) > 0 {
		I.Pop()
	}
}
