package main

// import "fmt"

type Node struct {
	val     int
	left    *Node
	right   *Node
	left_h  int
	right_h int
}

var BF int = 1

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func left_rotate(x, parent *Node) {

	y := x.right
	x.right = y.left
	y.left = x
	if parent != nil {
		if parent.left == x {
			parent.left = y
		} else {
			parent.right = y
		}
	}

	// Updating heights
	if x.right != nil {
		x.right_h = 1 + max(x.right.left_h, x.right.right_h)
	} else {
		x.right_h = 0
	}
	y.left_h = 1 + max(x.right_h, x.left_h)
}

func right_rotate(x, parent *Node) {

	y := x.left
	x.left = y.right
	y.right = x
	if parent != nil {
		if parent.left == x {
			parent.left = y
		} else {
			parent.right = y
		}
	}

	// Updating heights
	if x.left != nil {
		x.left_h = 1 + max(x.left.left_h, x.left.right_h)
	} else {
		x.left_h = 0
	}
	y.right_h = 1 + max(x.right_h, x.left_h)
}

func balance(node, parent *Node) {

	if node.left_h-node.right_h > 0 {
		// Left Heavy
		// fmt.Println("\tLeft-Heavy")
		l_h := node.left.left_h
		r_h := node.left.right_h
		if l_h >= r_h {
			// fmt.Println("\t\tUsing Right Rotate")
			right_rotate(node, parent)
		} else {
			// fmt.Println("\t\tUsing Left-Right Rotate")
			left_rotate(node.left, node)
			// fmt.Println("\t\t\tHalfway")
			right_rotate(node, parent)
		}
	} else {
		// Right Heavy
		// fmt.Println("\tRight-Heavy")
		l_h := node.right.left_h
		r_h := node.right.right_h
		if r_h >= l_h {
			// fmt.Println("\t\tUsing Left Rotate")
			left_rotate(node, parent)
		} else {
			// fmt.Println("\t\tUsing Right-Left Rotate")
			right_rotate(node.right, node)
			// fmt.Println("\t\t\tHalfway")
			left_rotate(node, parent)
		}
	}
}

func check_balance(node, parent *Node) {
	if node == nil {
		return
	}
	if abs(node.left_h-node.right_h) > BF {
		// fmt.Printf("DETECTED IMBALANCED AT %+v\n", node)
		balance(node, parent)
	}
}

func Insert(root, parent *Node, x int) (*Node, int) {
	if root == nil {
		new_Node := new(Node)
		new_Node.val = x
		return new_Node, 0
	}
	if root.val <= x {
		if root.right == nil {
			root.right_h += 1
			root.right, _ = Insert(root.right, root, x)
			check_balance(root, parent)
			return root.right, max(root.left_h, root.right_h)
		}
		new_node, r_h := Insert(root.right, root, x)
		root.right_h = r_h + 1
		check_balance(root, parent)
		return new_node, max(root.left_h, root.right_h)
	}
	if root.left == nil {
		root.left_h += 1
		root.left, _ = Insert(root.left, root, x)
		check_balance(root, parent)
		return root.left, max(root.left_h, root.right_h)
	}
	new_node, l_h := Insert(root.left, root, x)
	root.left_h = l_h + 1
	check_balance(root, parent)
	return new_node, max(root.left_h, root.right_h)
}

func Search(node *Node, x int) *Node {
	if node == nil {
		return nil
	} else if node.val == x {
		return node
	} else if x > node.val {
		return Search(node.right, x)
	}
	return Search(node.left, x)
}

func Update(node, parent *Node, x, y int) {
	Delete(node, parent, x)
	Insert(node, parent, y)
}

func find_leftmost(node *Node) int {
	if node.left == nil {
		return node.val
	}
	return find_leftmost(node.left)
}

func Delete(node, parent *Node, x int) *Node {

	var returned_node *Node

	// fmt.Printf("\tDELETE Func(%v, %v, %v)\n", node, parent, x)

	if node == nil {
		return nil
	}
	if node.val < x {
		returned_node = Delete(node.right, node, x)
		if returned_node != nil {
			node.right_h = 1 + max(returned_node.left_h, returned_node.right_h)
		} else {
			node.right_h = 0
		}
		// fmt.Printf("\t\tCHECKING = %v, %v\n", returned_node, node)
		check_balance(node, parent)
		// fmt.Printf("\tRETURN = %v\n", node)
		return node
	}
	if node.val > x {
		returned_node = Delete(node.left, node, x)
		if returned_node != nil {
			node.left_h = 1 + max(returned_node.left_h, returned_node.right_h)
		} else {
			node.left_h = 0
		}
		// fmt.Printf("\t\tCHECKING = %v, %v\n", returned_node, node)
		check_balance(node, parent)
		// fmt.Printf("\tRETURN = %v\n", node)
		return node
	}

	if node.right == nil {
		if parent.left == node {
			parent.left = node.left
			if node.left != nil {
				parent.left_h = max(node.left.left_h, node.left.right_h)
			} else {
				parent.left_h = 0
			}
		} else {
			parent.right = node.left
			if node.left != nil {
				parent.right_h = max(node.left.left_h, node.left.right_h)
			} else {
				parent.right_h = 0
			}
		}
		// fmt.Printf("\tRETURN = %v\n", node.left)
		return node.left
	} else {
		l_most_val := find_leftmost(node.right)
		node.val = l_most_val
		returned_node = Delete(node.right, node, l_most_val)
		if returned_node != nil {
			node.right_h = 1 + max(returned_node.left_h, returned_node.right_h)
		} else {
			node.right_h = 0
		}
		// fmt.Printf("\t\tCHECKING = %v, %v\n", returned_node, node)
		check_balance(node, parent)
		// fmt.Printf("\tRETURN = %v\n", node)
		return node
	}
}
