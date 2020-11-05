package segment

type Node struct {
	s, e        int
	left, right *Node
	sum         int
}

func BuildTree(arr []int, s, e int) *Node {
	root := &Node{s: s, e: e}
	if s == e {
		root.sum = arr[s]
		return root
	} else {
		m := (s + e) / 2
		root.left = BuildTree(arr, s, m)
		root.right = BuildTree(arr, m+1, e)
		root.sum = root.left.sum + root.right.sum
		return root
	}
}

func (n *Node) Sum(s, e int) int {
	if n == nil {
		return 0
	}
	if n.s >= s && n.e <= e {
		return n.sum
	}
	sum := 0
	m := (n.s + n.e) / 2
	if s <= m {
		sum += n.left.Sum(s, e)
	}
	if e > m {
		sum += n.right.Sum(s, e)
	}
	return sum
}
func (n *Node) Update(i int, val int) int {
	if n == nil {
		return 0
	}
	if n.s == i && n.e == i {
		diff := val - n.sum
		n.sum = val
		return diff
	}
	m := (n.s + n.e) / 2
	if i <= m {
		diff := n.left.Update(i, val)
		n.sum += diff
		return diff
	} else {
		diff := n.right.Update(i, val)
		n.sum += diff
		return diff
	}
}

// func main() {
// 	n := BuildTree([]int{1, 2, 3, 4}, 0, 3)
// 	fmt.Println(n.Sum(0, 3), n.Sum(1, 3), n.Sum(2, 2))
// }
