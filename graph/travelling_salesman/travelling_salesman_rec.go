package travellingsalesman

import "math"

func setAllNode(n int) uint64 {
	return uint64(1)<<n - 1
}

func FindDistRec(mat [][]int, node int, dest int, visited uint64) (int, bool) {
	visited = setNode(visited, node)
	if visited == setAllNode(len(mat)) {
		if cost := mat[node][dest]; cost != math.MaxInt32 {
			return cost, true
		}
	}

	var found bool
	var minDist int
	for next, cost := range mat[node] {
		if cost != math.MaxInt32 && !hasNode(visited, next) {
			dist, nFound := FindDistRec(mat, next, dest, visited)
			if nFound {
				dist += cost
				if found {
					if dist < minDist {
						minDist = dist
					}
				} else {
					minDist = dist
					found = true
				}
			}
		}
	}
	return minDist, found
}

func setNode(visited uint64, node int) uint64 {
	return visited | (uint64(1) << node)
}

func hasNode(visited uint64, node int) bool {
	return ((visited & (uint64(1) << node)) >> node) == 1
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

// 	fmt.Println(findDistRec(mat, 0, 0, 0))
// }
