package kahn

// TopologicalSort returns topologically sorted vertices for DAG and nil otherwise.
func TopologicalSort(adj [][]bool) []int {
	n := len(adj)

	inDegree := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				exist := adj[i][j]
				if exist {
					inDegree[j]++
				}
			}
		}
	}

	queue := make([]int, 0, n)
	for i, e := range inDegree {
		if e == 0 {
			queue = append(queue, i)
		}
	}
	var res []int
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		res = append(res, front)
		for i, exist := range adj[front] {
			if exist {
				inDegree[i]--
				if inDegree[i] == 0 {
					queue = append(queue, i)
				}
			}
		}
	}
	if len(res) == n {
		return res
	}
	// cycle
	return nil
}

// func main() {
// 	mat := [][]bool{
// 		0: {false, true, false, false},
// 		1: {false, false, true, false},
// 		2: {false, false, false, false},
// 		3: {true, false, true, false},
// 	}
// 	fmt.Println(topSort(mat))
// }
