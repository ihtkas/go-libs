package redblack

// import "fmt"

// func main() {
// 	var r *Node
// 	r = r.insert(1)
// 	r = r.insert(2)
// 	r = r.insert(3)
// 	r = r.insert(4)
// 	r = r.insert(6)
// 	// r = r.insert(7)
// 	// r = r.insert(8)
// 	// r = r.insert(9)
// 	print(r)
// }

// type Color int8

// const (
// 	Black Color = iota
// 	Red
// )

// func (c Color) String() string {
// 	if c == Black {
// 		return "B"
// 	}
// 	return "R"
// }

// type Node struct {
// 	val         int
// 	color       Color
// 	parent      *Node
// 	left, right *Node
// }

// func print(n *Node) {
// 	if n != nil {
// 		fmt.Printf("%s %d [%s %s]\n", n.color.String(), n.val, getKey(n.left), getKey(n.right))
// 		print(n.left)
// 		print(n.right)
// 	}
// }
// func getKey(n *Node) string {
// 	if n == nil {
// 		return "N"
// 	} else {
// 		return fmt.Sprint(n.val)
// 	}
// }
// func (root *Node) insert(k int) *Node {
// 	if root == nil {
// 		return &Node{val: k, color: Black}
// 	}
// 	n := &Node{val: k, color: Red}
// 	root.insertBST(n)
// 	return balance(root, n)
// }

// func (root *Node) insertBST(k *Node) {
// 	if root == nil {
// 		return
// 	} else if k.val < root.val {
// 		if root.left == nil {
// 			root.left = k
// 			k.parent = root
// 			return
// 		}
// 		root.left.insertBST(k)
// 	} else {
// 		if root.right == nil {
// 			root.right = k
// 			k.parent = root
// 			return
// 		}
// 		root.right.insertBST(k)
// 	}
// }

// func leftRotate(n *Node) {
// 	r := n.right
// 	n.right = r.left
// 	if n.right != nil {
// 		n.right.parent = n
// 	}
// 	r.left = n
// 	r.parent = n.parent
// 	if n.parent != nil {
// 		if n.parent.left == n {
// 			n.parent.left = r
// 		} else {
// 			n.parent.right = r
// 		}
// 	}
// 	n.parent = r
// }

// func rightRotate(n *Node) {
// 	l := n.left
// 	n.left = l.right
// 	if n.left != nil {
// 		n.left.parent = n
// 	}
// 	l.right = n
// 	l.parent = n.parent
// 	if n.parent != nil {
// 		if n.parent.left == n {
// 			n.parent.left = l
// 		} else {
// 			n.parent.right = l
// 		}
// 	}
// 	n.parent = l
// }

// func balance(root *Node, n *Node) *Node {
// 	for n.parent != nil && n.color == Red && n.parent.color == Red {
// 		p := n.parent
// 		gp := n.parent.parent
// 		var pSib *Node
// 		isPLeft := gp.left == p

// 		if isPLeft {
// 			pSib = gp.right
// 			isPLeft = true
// 		} else {
// 			pSib = gp.left
// 		}

// 		if pSib != nil && pSib.color == Red {
// 			pSib.color = Black
// 			p.color = Black
// 			gp.color = Red
// 			p = gp
// 		} else {
// 			isLeft := p.left == n
// 			switch {
// 			case isLeft && isPLeft:
// 				rightRotate(gp)
// 				gp.color, p.color = p.color, gp.color
// 				n = p
// 			case !isLeft && isPLeft:
// 				leftRotate(p)
// 				rightRotate(gp)
// 				p = n
// 				gp.color, p.color = p.color, gp.color
// 			case !isLeft && !isPLeft:
// 				fmt.Println("here")
// 				print(root)
// 				leftRotate(gp)

// 				gp.color, p.color = p.color, gp.color
// 				n = p
// 				fmt.Println("here---")
// 				print(p)
// 				fmt.Println("here---")
// 			case isLeft && !isPLeft:
// 				rightRotate(p)
// 				leftRotate(gp)
// 				p = n
// 				gp.color, p.color = p.color, gp.color
// 			}
// 		}
// 	}

// 	if n.parent == nil {
// 		root = n
// 	}
// 	root.color = Black
// 	return root
// }
