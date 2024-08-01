package main

import "fmt"

func postorder(node *Node) {
	if node != nil {
		fmt.Printf("%p => \t%+v\n", node, node)
		postorder(node.left)
		postorder(node.right)
	}
}

func main() {
	query := []int{7, -2, 8, 2, -1, 3, 4, 0, 10, -10, 20, -20, 15, -15}

	root := new(Node)

	n := len((query))
	for i := 0; i < n; i++ {
		if i == 0 {
			root.right, _ = Insert(nil, root, query[i])
		} else {
			Insert(root.right, root, query[i])
		}
	}

	postorder(root.right)
	fmt.Println()
	Delete(root.right, root, 7)
	fmt.Println()
	postorder(root.right)
}
