// Package AVL implements the AVL (Adelson-Velskii and Landis) self-balancing
// binary search tree.
package AVL

// Any value type that satisifies AVL.Comparable interface
// can be managed by an AVL data structure
type Comparable interface {
	Equal(Comparable) bool
	Less(Comparable) bool
}

type Node struct {
	H       int
	V       Comparable
	p, l, r *Node
}

func NewNode(v Comparable) *Node {
	return &Node{H: 1, V: v} // p, l, r == nil
}

func nht(n *Node) int {
	if n != nil {
		return n.H
	} else {
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Calculate the height of the node
func (b *Node) HCalc() {
	b.H = max(nht(b.l), nht(b.r)) + 1
}

// Rotate b Left
func (b *Node) LRot() {

	rc := b.r
	c := b.r.l
	b.r.l = b
	b.r = c
	if c != nil {
		c.p = b.r
	}
	if b.p != nil {
		if b == b.p.l {
			// left child of parent
			b.p.l = rc
		} else if b == b.p.r {
			// right child of parent
			b.p.r = rc
		} else {
			panic("LRot: impossible")
		}
	}
	rc.p, b.p = b.p, rc

	b.HCalc()
	rc.HCalc()
}

// Rotate b Right
func (b *Node) RRot() {
	lc := b.l
	c := b.l.r
	b.l.r = b
	b.l = c
	if c != nil {
		c.p = b.l
	}
	if b.p != nil {
		if b == b.p.l {
			b.p.l = lc
		} else if b == b.p.r {
			b.p.r = lc
		} else {
			panic("RRot: impossible")
		}
	}
	lc.p, b.p = b.p, lc

	b.HCalc()
	lc.HCalc()
}

// Fix the node imbalance going up the tree, return the new root
func (b *Node) Fix() (root *Node) {
	// update the depth of elements in the path to get here
	root = b
	node := b
	for node != nil {
		node.HCalc()
		hl, hr := nht(node.l), nht(node.r)
		if hr > 1+hl {
			chl, chr := nht(node.r.l), nht(node.r.r)
			if chr >= chl {
				node.LRot()
			} else {
				node.r.RRot()
				node.LRot()
			}
		} else if hl > 1+hr {
			chl, chr := nht(node.l.l), nht(node.l.r)
			if chr <= chl {
				node.RRot()
			} else {
				node.l.LRot()
				node.RRot()
			}
		}
		root, node = node, node.p
	}
	return
}

// return node and parent
// returning a nil parent means it's the node itself
func (b *Node) Find(v Comparable) (c, p *Node) {
	c, p = b, b.p
	for c != nil {
		if v.Equal(c.V) {
			break
		}

		if v.Less(c.V) {
			c, p = c.l, c
		} else {
			c, p = c.r, c
		}
	}
	return
}

// Insert a node at the right place; return the new
// root.
func (b *Node) Insert(v Comparable) *Node {
	c, p := b.Find(v)
	if c != nil {
		return b
	}

	if v.Less(p.V) {
		p.l = &Node{p: p, V: v, H: 1}
	} else {
		p.r = &Node{p: p, V: v, H: 1}
	}

	return p.Fix()
}

// return the smallest value
func (b *Node) Smallest() Comparable {
	x := b
	for x.l != nil {
		x = x.l
	}
	return x.V
}

// return the largest value
func (b *Node) Largest() Comparable {
	x := b
	for x.r != nil {
		x = x.r
	}
	return x.V
}

// Remove a node, rebalance and return the new root
func (b *Node) Remove(v Comparable) *Node {
	c, p := b.Find(v)
	if c == nil {
		return nil
	}

	if c.l == nil {
		// no left child; replace node with it's right child
		if p == nil {
			// c is at the root
			if c.r != nil {
				c.r.p = nil
				return c.r // c.r wasn't touched so no need to Fix
			}
			// c was the last element
			return nil
		}

		if c == p.l {
			// c is left child of p
			p.l = c.r
		} else {
			p.r = c.r
		}
		if c.r != nil {
			c.r.p = p
		}
		return p.Fix()
	} else {
		// the rightmost leaf of the left child (highest value on
		// the left side) can asecend to this node's position. by
		// definition it wont have a right child
		// r is the replacement, rp is its current parent
		r, rp := c.l, c
		for r.r != nil {
			r, rp = r.r, r
		}

		r.p = p
		if p != nil {
			if c == p.l {
				p.l = r
			} else {
				p.r = r
			}
		}
		r.r = c.r
		if c.r != nil {
			c.r.p = r
		}
		if r != c.l {
			// not the left node of the node being replaced.
			rp.r = r.l
			return rp.Fix()
		}
		return r.Fix()
	}
}

// Traverse the tree in ascending order (as determined by to Comparable.Less)
func (b *Node) Traverse(f func(interface{})) {
	if b.l != nil {
		b.l.Traverse(f)
	}
	f(b.V)
	if b.r != nil {
		b.r.Traverse(f)
	}
}

// Traverse the tree in descending order (as determined by Comparable.Less)
func (b *Node) RTraverse(f func(interface{})) {
	if b.r != nil {
		b.r.RTraverse(f)
	}
	f(b.V)
	if b.l != nil {
		b.l.RTraverse(f)
	}
}
