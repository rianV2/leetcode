package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Print() {
	fmt.Println(n.Val)
	if n.Left != nil {
		n.Left.Print()
	}
	if n.Right != nil {
		n.Right.Print()
	}
}

func NewTreeFromArray(inp []int) *TreeNode {
	totalInp := len(inp)
	if totalInp == 0 {
		return nil
	}

	var queue []*TreeNode

	root := &TreeNode{Val: inp[0]}
	queue = append(queue, root)

	i := 1

	for len(queue) > 0 && i < totalInp {
		node := queue[0]
		queue = queue[1:] // remove from queue

		// Process left child
		if i < totalInp && inp[i] != -1 {
			// Create node with & operator directly, not as local variable
			node.Left = &TreeNode{Val: inp[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Process right child
		if i < totalInp && inp[i] != -1 {
			// Create node with & operator directly, not as local variable
			node.Right = &TreeNode{Val: inp[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	lv := 0
	for len(queue) > 0 {
		current := len(queue)

		for i := 0; i < current; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		lv++
	}

	return lv
}

func main() {
	root := []int{3, 9, 20, -1, -1, 15, 7}

	fmt.Println(root)
	rootTree := NewTreeFromArray(root)
	fmt.Println("---print---")
	rootTree.Print()
	fmt.Println("---")
	maxDepth := maxDepth(rootTree)
	fmt.Println(maxDepth)
}
