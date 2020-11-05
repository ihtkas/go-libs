package floydwarshall

import "math"

const m = int(math.MaxUint32 >> 1)

func AllShortestPath(adj [][]int) ([][]int, [][]int) {
	n := len(adj)
	path := make([][]int, n)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		path[i] = make([]int, n)
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if adj[i][j] == 0 || adj[i][j] == m {
				path[i][j] = n
			} else {
				path[i][j] = i
			}
			dist[i][j] = adj[i][j]
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				nd := dist[i][k] + dist[k][j]
				if dist[i][k] != m && dist[k][j] != m && dist[i][j] > nd {
					dist[i][j] = nd
					path[i][j] = path[k][j]
				}
			}
		}
	}
	return dist, path
}

// func main() {
// 	adj := [][]int{
// 		0: {0, 3, 6, 15},
// 		1: {m, 0, -2, m},
// 		2: {m, m, 0, 2},
// 		3: {1, m, m, 0},
// 	}

// 	fmt.Println(allShortestPath(adj))
// }
