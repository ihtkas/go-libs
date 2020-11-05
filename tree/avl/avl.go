package avl

import (
	"fmt"
)

type Node struct {
	key         int
	left, right *Node
	height      int
}

func print(n *Node) {
	if n != nil {
		fmt.Printf("%d [%s %s]\n", n.key, getKey(n.left), getKey(n.right))
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

// Insert will insert the key in to tree. root can be nil.
func Insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key, height: 1}
	}
	if key < root.key {
		root.left = Insert(root.left, key)
	} else if key > root.key {
		root.right = Insert(root.right, key)
	} else {
		return root
	}
	bal := height(root.left) - height(root.right)

	if bal > 1 && key < root.left.key {
		// left left
		return rightRotate(root)
	} else if bal > 1 && key > root.left.key {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	} else if bal < -1 && key > root.right.key {
		return leftRotate(root)
	} else if bal < -1 && key < root.right.key {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	} else {
		root.height = 1 + max(height(root.left), height(root.right))
		return root
	}

}

// Delete will remove the node with given key
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
			fmt.Println("max", key, m, root.key)
			root.key = m

			root.left = Delete(root.left, root.key)
		} else if root.right != nil {
			root.key = root.right.key
			root.right = Delete(root.right, root.right.key)
		} else {
			return nil
		}
	}
	bal := height(root.left) - height(root.right)

	if bal > 1 && height(root.left.left) > height(root.left.right) {
		// left left
		return rightRotate(root)
	} else if bal > 1 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	} else if bal < -1 && height(root.right.right) > height(root.right.left) {
		return leftRotate(root)
	} else if bal < -1 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	} else {
		root.height = 1 + max(height(root.left), height(root.right))
		return root
	}

}

func leftRotate(node *Node) *Node {
	r := node.right
	node.right = r.left
	r.left = node
	node.height = 1 + max(height(node.left), height(node.right))
	r.height = 1 + max(height(r.left), height(r.right))
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
	node.left = l.right
	l.right = node
	node.height = 1 + max(height(node.left), height(node.right))
	l.height = 1 + max(height(l.left), height(l.right))

	return l
}

func height(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

// func main() {
// 	var root *Node
// 	root = Insert(root, 10)
// 	root = Insert(root, 5)
// 	root = Insert(root, 4)
// 	root = Insert(root, 1)
// 	root = Insert(root, 6)
// 	root = Insert(root, 8)
// 	root = Insert(root, 7)
// 	root = Insert(root, 9)
// 	root = Insert(root, 2)
// 	root = Insert(root, 3)
// 	root = Delete(root, 7)
// 	root = Delete(root, 6)
// 	root = Delete(root, 8)
// 	root = Delete(root, 1)
// 	root = Delete(root, 5)
// 	print(root)
// 	root = Delete(root, 2)
// 	print(root)
// 	root = Delete(root, 4)
// 	print(root)
// }
