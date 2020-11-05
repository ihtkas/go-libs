package dijkstra

import "fmt"

// ShortestPathAndDistance returns the nodes between src and dest in the shortest path and the shortest distance.
func ShortestPathAndDistance(adjList [][][]int, src, dest int) ([]int, int) {
	heap := newMinHeap()
	heap.add(&Node{k: src, v: 0})
	dist := make([]int, len(adjList))
	path := make([]int, len(adjList))
	// fmt.Println(adjList)
	for !heap.isEmpty() {
		// fmt.Println(heap)
		curr := heap.extractMin()
		dist[curr.k] = curr.v

		if curr.k == dest {
			break
		}
		for _, e := range adjList[curr.k] {
			if n := heap.search(e[0]); n != nil {
				if curr.v+e[1] < n.v {
					heap.decrease(n, curr.v+e[1])
					path[e[0]] = curr.k
				}
			} else {
				heap.add(&Node{k: e[0], v: curr.v + e[1]})
				path[e[0]] = curr.k
			}
		}
	}
	var shortPath []int
	if dist[dest] != 0 {
		root := path[dest]
		for root != src {
			shortPath = append(shortPath, root)
			root = path[root]
		}
		j := len(shortPath) - 1
		// reverese the path
		for i := 0; i < j; i++ {
			shortPath[i], shortPath[j] = shortPath[j], shortPath[i]
			j--
		}
	}
	return shortPath, dist[dest]
}

type MinHeap struct {
	nodes []*Node
	pos   map[int]int
}

func (h *MinHeap) String() string {
	return fmt.Sprint(h.nodes)
}

func newMinHeap() *MinHeap {
	return &MinHeap{pos: make(map[int]int)}
}

type Node struct {
	k, v int
}

func (n *Node) String() string { return fmt.Sprintf("%d:%d", n.k, n.v) }

func (h *MinHeap) isEmpty() bool {
	return len(h.nodes) == 0
}

func (h *MinHeap) add(n *Node) {
	h.nodes = append(h.nodes, n)
	currInd := len(h.nodes) - 1
	h.pos[currInd] = currInd

	pInd := (currInd+1)/2 - 1

	for pInd <= 0 {
		parent := h.nodes[pInd]
		if n.v < parent.v {
			h.nodes[currInd], h.nodes[pInd] = h.nodes[pInd], h.nodes[currInd]
			h.pos[parent.k] = currInd
			h.pos[n.k] = pInd
			currInd = pInd
			pInd = (pInd+1)/2 - 1
		} else {
			break
		}
	}
}

func (h *MinHeap) search(k int) *Node {
	if ind, exist := h.pos[k]; exist {
		return h.nodes[ind]
	}
	return nil
}

// decrease assumes that newVal will be lesser than existing value in Node n
func (h *MinHeap) decrease(n *Node, newVal int) {
	currInd := h.pos[n.k]
	n.v = newVal
	pInd := (currInd+1)/2 - 1

	for pInd <= 0 {
		parent := h.nodes[pInd]
		if n.v < parent.v {
			h.nodes[currInd], h.nodes[pInd] = h.nodes[pInd], h.nodes[currInd]
			h.pos[parent.k] = currInd
			h.pos[n.k] = pInd
			currInd = pInd
			pInd = (pInd+1)/2 - 1
		} else {
			break
		}
	}
}

func (h *MinHeap) extractMin() *Node {
	size := len(h.nodes) - 1
	min := h.nodes[0]
	h.nodes[0] = h.nodes[size]
	h.pos[h.nodes[0].k] = 0
	delete(h.pos, min.k)
	h.nodes = h.nodes[:size]
	i := 0
	for i < size {
		min := i
		c1 := i*2 + 1
		c2 := i*2 + 1
		if c1 >= size {
			break
		}
		if h.nodes[c1].v < h.nodes[i].v {
			min = c1
		}
		if c2 < size && h.nodes[c1].v < h.nodes[min].v {
			min = c2
		}

		if i == min {
			break
		} else {
			h.pos[h.nodes[min].k] = i
			h.pos[h.nodes[i].k] = min
			h.nodes[min], h.nodes[i] = h.nodes[i], h.nodes[min]
			// now min pos has changed and this has to be checked with its children.
			i = min
		}
	}
	return min
}

// func main() {
// 	adjList := [][][]int{
// 		0: {{1, 1}},
// 		1: {{2, 1}, {3, 2}},
// 		2: {{4, 3}, {5, 8}},
// 		3: {{5, 10}},
// 		4: {{5, 3}},
// 		5: {},
// 	}
// 	fmt.Println(ShortestPathAndDistance(adjList, 0, 5))
// }
