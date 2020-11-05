package kruskal

import "fmt"

func main() {
	edges := [][]int{{0, 1, 28}, {1, 2, 16}, {2, 3, 12}, {3, 4, 22}, {4, 5, 25}, {4, 6, 24},
		{5, 0, 10}, {3, 6, 18}, {6, 1, 14}}
	fmt.Println(minSpanningTree(edges, 7))
}

// minSpanningTree assumes input edges forms a connected for given number of nodes n (0 to n-1 nodes).
// Returns edges forming min spanning tree and min weight
func minSpanningTree(edges [][]int, n int) ([][]int, int) {
	quickSort(edges, 0, len(edges)-1)
	s := newDSU(n)
	var weight int
	var spanningTree [][]int
	for _, edge := range edges {
		if s.union(edge[0], edge[1]) {
			spanningTree = append(spanningTree, edge)
			weight += edge[2]
		}

	}
	return spanningTree, weight
}

func quickSort(arr [][]int, lo, hi int) {
	if lo < hi {
		p := partition(arr, lo, hi)
		quickSort(arr, lo, p-1)
		quickSort(arr, p+1, hi)
	}
}

func partition(arr [][]int, lo, hi int) int {
	piv := arr[hi][2]
	i := lo
	j := hi - 1
	for i <= j {
		if arr[j][2] < piv {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		} else {
			j--
		}
	}
	j++
	arr[j], arr[hi] = arr[hi], arr[j]
	return j
}

// disjoint set union
type dsu []int

func newDSU(n int) dsu {
	s := make([]int, n)
	for i := range s {
		s[i] = -1
	}
	return s
}

func (s dsu) find(i int) int {
	for s[i] >= 0 {
		i = s[i]
	}
	return i
}

// union joins only if i and j belongs to different set or else returns false
func (s dsu) union(i, j int) bool {
	fmt.Println(s)
	iRef := s.find(i)
	jRef := s.find(j)
	// Wrong impl. Need to change
	if iRef == jRef {
		return false
	}
	if s[iRef] < s[jRef] {
		s[iRef] += s[jRef]
		s[jRef] = iRef
	} else {
		s[jRef] += s[iRef]
		s[iRef] = jRef
	}
	return true
}
