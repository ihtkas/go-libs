package redblack

import (
	"fmt"
)

type Color int8

const (
	Red   Color = 1
	Black Color = 2
)

func (c Color) String() string {
	if c == Red {
		return "R"
	} else {
		return "B"
	}
}

type Node struct {
	key         int
	left, right *Node
	color       Color
	parent      *Node
}

func print(n *Node) {
	if n != nil {
		fmt.Printf("%s %d [%s %s]\n", n.color, n.key, getKey(n.left), getKey(n.right))
		print(n.left)
		print(n.right)
	}
}
func getKey(n *Node) string {
	if n == nil {
		return "N"
	} else {
		return fmt.Sprint(n.key)
	}
}

func maxKey(node *Node) int {
	if node.right != nil {
		return maxKey(node.right)
	}
	return node.key
}

func Insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key, color: Black}
	}

	node := &Node{key: key, color: Red}
	insert(root, node)
	fmt.Println("after bst")
	print(root)
	root = colorBalance(node)
	return root
}

func colorBalance(node *Node) *Node {
	r := findRoot(node)
	for node.parent != nil {
		// fmt.Println("node", node.key, node.parent.color)
		p := node.parent
		if p.color == Black {
			return findRoot(p)
		}
		// can't be nil as parent is red and there will be atleast root which is black
		gp := p.parent

		isPLeft := gp.left == p

		var pSibling *Node
		if isPLeft {
			pSibling = gp.right
		} else {
			pSibling = gp.left
		}

		var isPSiblingRed bool

		if pSibling != nil {
			isPSiblingRed = pSibling.color == Red
		}
		// fmt.Println("node", node.key, node.parent.color)
		if isPSiblingRed {
			gp.color = Red
			if pSibling != nil {
				pSibling.color = Black
			}
			p.color = Black
			node = gp
		} else {
			isNodeLeft := p.left == node
			fmt.Println("node", node.key, isNodeLeft, isPLeft)
			switch {
			case isNodeLeft && isPLeft: // LL
				node = rightRotate(gp)

			case !isNodeLeft && isPLeft: // RL
				leftRotate(p)
				print(r)
				fmt.Println("----")
				node = rightRotate(gp)
			case !isNodeLeft && !isPLeft: // RR
				node = leftRotate(gp)
			default: // LR
				rightRotate(p)
				node = leftRotate(gp)
			}
		}
	}
	node.color = Black
	return node
}

func findRoot(node *Node) *Node {
	for node.parent != nil {
		node = node.parent
	}
	return node
}

func insert(root *Node, key *Node) {
	if root == nil {
		return
	}
	if key.key < root.key {
		if root.left == nil {
			root.left = key
			key.parent = root
			return
		}
		insert(root.left, key)
	} else if key.key > root.key {
		if root.right == nil {
			root.right = key
			key.parent = root
			return
		}
		insert(root.right, key)
	}

}

func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = Delete(root.left, key)
	} else if key > root.key {
		root.right = Delete(root.right, key)
	} else {
		if root.left != nil {
			m := maxKey(root.left)
			// fmt.Println("max", key, m, root.key)
			root.key = m
			root.left = Delete(root.left, root.key)
		} else if root.right != nil {
			root.key = root.right.key
			root.right = Delete(root.right, root.right.key)
		} else {
			return nil
		}
	}
	return root

}

func leftRotate(node *Node) *Node {
	r := node.right
	r.parent = node.parent
	if node.parent != nil {
		if node == node.parent.left {
			node.parent.left = r
		} else {
			node.parent.right = r
		}
	}
	node.right = r.left
	if r.left != nil {
		r.left.parent = node
	}
	r.left = node
	node.parent = r
	node.color, r.color = r.color, node.color
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rightRotate(node *Node) *Node {
	l := node.left
	l.parent = node.parent
	if node.parent != nil {
		if node == node.parent.left {
			node.parent.left = l
		} else {
			node.parent.right = l
		}
	}
	node.left = l.right
	if l.right != nil {
		l.right.parent = node
	}
	l.right = node
	node.parent = l
	node.color, l.color = l.color, node.color
	return l
}

// func height(node *Node) int {
// 	if node == nil {
// 		return 0
// 	} else {
// 		return node.height
// 	}
// }

// func main() {
// 	var root *Node
// 	// root = insert(root, 10)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 5)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 4)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 1)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 6)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 8)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 7)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 9)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 2)
// 	// fmt.Println()
// 	// print(root)
// 	// root = insert(root, 3)
// 	root = insert(root, 10)
// 	root = insert(root, 20)
// 	root = insert(root, -10)
// 	root = insert(root, 15)
// 	root = insert(root, 17)
// 	root = insert(root, 40)
// 	root = insert(root, 50)
// 	root = insert(root, 60)

// 	fmt.Println()
// 	print(root)
// 	// root = Delete(root, 7)
// 	// root = Delete(root, 6)
// 	// root = Delete(root, 8)
// 	// root = Delete(root, 1)
// 	// root = Delete(root, 5)
// 	// print(root)
// 	// root = Delete(root, 2)
// 	// print(root)
// 	// root = Delete(root, 4)
// 	// print(root)
// }
