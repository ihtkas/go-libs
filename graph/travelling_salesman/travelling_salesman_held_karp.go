package travellingsalesman

import (
	"fmt"
	"math"
)

type Index struct {
	dest    int
	setHash uint64
}

func FindDistHeldKarp(mat [][]int) ([]int, int) {
	n := len(mat) - 1
	setRefs := generateSets(n)
	sets := make([]*Set, len(setRefs))
	copy(sets, setRefs)
	quickSort(sets)

	costs := make(map[Index]int)
	predecessors := make(map[Index]int)
	for i := 1; i <= n; i++ {
		ind := Index{i, 0}
		costs[ind] = mat[0][i]
		predecessors[ind] = 0
	}
	for i := 1; i < len(sets)-1; i++ {
		nodes := sets[i]
		// fmt.Println("processing", nodes)
		for j := 1; j <= n; j++ {
			if !nodes.arr.has(j) {

				minCost := int(math.MaxInt32)
				var predecessor int
				for _, node := range nodes.arr {
					rem := remove(nodes, node, setRefs)
					cost := mat[node][j] + costs[Index{node, rem.hash}]
					// fmt.Println("compare", mat[node][j], costs[Index{node, rem.hash}], rem)
					if cost < minCost {
						minCost = cost
						predecessor = node
					}
				}
				ind := Index{j, nodes.hash}
				costs[ind] = minCost
				predecessors[ind] = predecessor
				// fmt.Println("---", nodes, "s:", predecessor, "min:", minCost)
			}

		}

	}

	minCost := int(math.MaxInt32)
	var predecessor int
	nodes := sets[len(sets)-1]
	for _, node := range nodes.arr {
		rem := remove(nodes, node, setRefs)
		cost := mat[node][0] + costs[Index{node, rem.hash}]
		// fmt.Println("compare", mat[node][0], costs[Index{node, rem.hash}], rem)
		if cost < minCost {
			minCost = cost
			predecessor = node
		}
	}
	ind := Index{0, nodes.hash}
	costs[ind] = minCost
	predecessors[ind] = predecessor
	// fmt.Println("---", nodes, "s:", predecessor, "min:", minCost)
	// fmt.Println(predecessor)
	path := []int{0, predecessor}
	for predecessor != 0 {
		rem := remove(nodes, predecessor, setRefs)
		predecessor = predecessors[Index{predecessor, rem.hash}]
		nodes = rem
		path = append(path, predecessor)
		// fmt.Println(predecessor)
	}

	return path, minCost
}

func generateSets(n int) []*Set {
	num := uint64(1 << n)

	res := make([]*Set, num)
	for i := uint64(0); i < num; i++ {
		s := &Set{hash: i}
		comb := i
		for j := 0; j < n; j++ {

			if comb&1 == 1 {
				s.add(j + 1)
			}
			comb = comb >> 1
		}
		res[i] = s
	}

	return res
}

type Set struct {
	hash uint64
	arr  OrdSet
}

func (s *Set) String() string { return fmt.Sprintf("%d:%s", s.hash, fmt.Sprint(s.arr)) }

func (s *Set) add(k int) {
	s.arr = s.arr.add(k)
}

func remove(s *Set, k int, setRef []*Set) *Set {
	n := uint64(1) << (k - 1)
	n = n ^ math.MaxUint64
	hash := s.hash & n
	return setRef[hash]
}

type OrdSet []int

func (s OrdSet) equal(s2 OrdSet) bool {
	if len(s) != len(s2) {
		return false
	}
	for i, e := range s {
		if e != s2[i] {
			return false
		}
	}
	return true

}

func (s OrdSet) has(k int) bool {
	_, found := s.find(k)
	return found
}

func (s OrdSet) add(k int) OrdSet {

	pos, found := s.find(k)
	if found {
		return s
	} else {
		res := make([]int, len(s)+1)
		res[pos] = k
		copy(res, s[:pos])
		copy(res[pos+1:], s[pos:])
		return res
	}

}

func (s OrdSet) find(k int) (int, bool) {
	lo := 0
	hi := len(s)

	for lo < hi {
		mid := (lo + hi) / 2
		if s[mid] == k {
			return mid, true
		} else if k < s[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo, false
}

// func clone(s OrdSet) OrdSet {
// 	m := make(map[int]struct{})
// 	for node, _ := range s {
// 		m[node] = struct{}{}
// 	}
// 	return m
// }
// TODO: This is a an exercise. Replace sort with go's sort lib
func quickSort(arr []*Set) {
	quickSortRec(arr, 0, len(arr)-1)
}

func quickSortRec(arr []*Set, s, e int) {
	if s < e {
		p := partition(arr, s, e)
		quickSortRec(arr, s, p-1)
		quickSortRec(arr, p+1, e)
	}
}

func partition(arr []*Set, s, e int) int {
	p := len(arr[e].arr)
	i := s
	j := e - 1
	for i <= j {
		if len(arr[j].arr) < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		} else {
			j--
		}
	}
	j++
	arr[j], arr[e] = arr[e], arr[j]
	return j
}

// func main() {
// 	mx := int(math.MaxInt32)
// 	// mat := [][]int{
// 	// 	0: {mx, 10, 15, 20},
// 	// 	1: {5, mx, 9, 10},
// 	// 	2: {6, 13, mx, 12},
// 	// 	3: {8, 8, 9, mx},
// 	// }
// 	// mat := [][]int{
// 	// 	0: {mx, 1, 15, 6},
// 	// 	1: {2, mx, 7, 3},
// 	// 	2: {9, 6, mx, 12},
// 	// 	3: {10, 4, 8, mx},
// 	// }

// 	// mat := [][]int{
// 	// 	0: {mx, 10, mx, mx},
// 	// 	1: {mx, mx, 9, mx},
// 	// 	2: {mx, mx, mx, 12},
// 	// 	3: {8, mx, mx, mx},
// 	// }

// 	mat := [][]int{
// 		0: {mx, 10, 2, mx},
// 		1: {mx, mx, 9, 2},
// 		2: {mx, 1, mx, 12},
// 		3: {8, mx, mx, mx},
// 	}
// 	fmt.Println(findMinDistPath(mat))
// }
