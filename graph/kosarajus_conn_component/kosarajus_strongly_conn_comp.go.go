package kosaraju

func FindStrongComponents(adj [][]bool) [][]int {
	var finishStack []int
	visited := make([]bool, len(adj))
	for n := range adj {
		buildFinishStack(adj, n, &finishStack, visited)
	}
	visited = make([]bool, len(adj))
	l := len(adj)
	revAdj := make([][]bool, l)
	for i := range revAdj {
		revAdj[i] = make([]bool, l)
	}

	for src, list := range adj {
		for dest, exist := range list {
			revAdj[dest][src] = exist
		}
	}
	var res [][]int
	top := l - 1
	for top >= 0 {
		comps := buildConnectedComps(revAdj, finishStack[top], visited)
		if len(comps) != 0 {
			res = append(res, comps)
		}
		top--

	}
	return res
}

func buildFinishStack(adj [][]bool, node int, finishStackP *[]int, visited []bool) {
	if !visited[node] {
		visited[node] = true
		for n, exist := range adj[node] {
			if exist {
				if !visited[n] {
					buildFinishStack(adj, n, finishStackP, visited)
				}
			}
		}
		*finishStackP = append((*finishStackP), node)
	}
}

func buildConnectedComps(adj [][]bool, node int, visited []bool) []int {

	if !visited[node] {
		visited[node] = true
		connectedComps := []int{node}

		children := ""
		for n, exist := range adj[node] {
			if exist {
				children += string('a'+n) + " "
				if !visited[n] {
					connectedComps = append(connectedComps, buildConnectedComps(adj, n, visited)...)
				}
			}
		}
		// fmt.Println(string('a'+node), ":", children)
		return connectedComps
	}
	return nil
}

const (
	a = 0
	b = 1
	c = 2
	d = 3
	e = 4
	f = 5
	g = 6
	h = 7
	i = 8
	j = 9
	k = 10
)

// func main() {

// 	adjList := [][]int{
// 		a: {b},
// 		b: {c, d},
// 		c: {a},
// 		d: {e},
// 		e: {f},
// 		f: {d},
// 		g: {f, h},
// 		h: {i},
// 		i: {j},
// 		j: {g, k},
// 		k: {},
// 	}
// 	adj := make([][]bool, 11)
// 	for i := range adj {
// 		adj[i] = make([]bool, 11)
// 	}
// 	for src, list := range adjList {
// 		for _, dest := range list {
// 			adj[src][dest] = true
// 		}
// 	}
// 	for _, comps := range FindStrongComponents(adj) {
// 		str := ""
// 		for _, node := range comps {
// 			str += string('a' + node)
// 		}
// 		fmt.Println(str)
// 	}
// }
