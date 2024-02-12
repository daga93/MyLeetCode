package main

type Node struct {
	Candy int
	Left  *Node
	Right *Node
}

func main() {
	// build tree
	n1 := newHouse(4)
	n2 := newHouse(9)
	b := newNonHouse(n1, n2)
	n4 := newHouse(15)
	c := newNonHouse(b, n4)

}

func newHouse(candy int) *Node {
	return &Node{Candy: candy}
}

func newNonHouse(left, right *Node) *Node {
	return &Node{Left: left, Right: right}
}

func (n *Node) treeCandy() int {
	curCandies := 0

	if n.Left != nil && n.Right != nil {
		curCandies = n.Left.treeCandy() + n.Right.treeCandy()
	} else {
		curCandies = n.Candy
	}
	return curCandies

}
