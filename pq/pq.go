// heap based priority queue. the standard container/heap is better.
// this is an exercise
package PQ

// import "fmt"

type Node struct {
	P int
	V interface{}
}

type PQ []Node

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

// make sure parent's priority >= max(left,right) children's priority
func (I *PQ) downheap(i int) {
	for i >= 0 && i < len(*I) {
		l := left(i)
		r := right(i)
		var largest int
		if l < len(*I) && (*I)[l].P > (*I)[i].P {
			largest = l
		} else {
			largest = i
		}
		if r < len(*I) && (*I)[r].P > (*I)[largest].P {
			largest = r
		}

		if largest == i {
			break
		}

		// fmt.Println("I[", i, "] ↔ I[", largest, "]")
		// fmt.Println((*I)[i], "↔", (*I)[largest])
		(*I)[i], (*I)[largest] = (*I)[largest], (*I)[i]
		i = largest // essentially I.downheap(largest)
	}
}

// Convert a []Node to a PQ
func (I *PQ) MakePQ() {
	// fmt.Println("MakePQ:", *I)
	for i := (len(*I) - 1) / 2; i >= 0; i-- {
		// fmt.Println("MakePQ i: ", i)
		I.downheap(i)
	}
}

// Return the maximum value
func (I *PQ) Pop() (val Node) {
	val = (*I)[0]
	(*I)[0] = (*I)[len(*I)-1]
	*I = (*I)[:len(*I)-1]
	I.downheap(0)
	return
}

// make sure everything up the heap from here obeys the priority
// ordering -- which says parent's priority >= max(left,right) child priority
func (I *PQ) upheap(i int) {
	ip := parent(i)
	for ip >= 0 && (*I)[i].P > (*I)[ip].P {
		(*I)[i], (*I)[ip] = (*I)[ip], (*I)[i]
		i, ip = ip, parent(ip)
	}
}

// add a new element, maintaining the heap structure.
func (I *PQ) Push(p int, v interface{}) {
	*I = append(*I, Node{P: p, V: v})
	I.upheap(len(*I) - 1)
}
