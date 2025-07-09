package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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

	// list of node we need to fill in
	queue = append(queue, root)

	i := 1

	for len(queue) > 0 && i < totalInp {
		node := queue[0]
		queue = queue[1:] //remove from queue
		// fmt.Println(queue)

		if i < totalInp && inp[i] != -1 {
			nodeLeft := TreeNode{
				Val: inp[i],
			}
			queue = append(queue, &nodeLeft)
			node.Left = &nodeLeft
			// fmt.Println("inp", inp[i])
		}
		i++

		if i < totalInp && inp[i] != -1 {
			nodeRight := TreeNode{
				Val: inp[i],
			}
			queue = append(queue, &nodeRight)
			node.Right = &nodeRight
			// fmt.Println("inp", inp[i])
		}
		i++
		if node.Left != nil {
			// fmt.Println("nodeLeft", node.Left.Val)
		}
		// fmt.Println("--debug")
	}

	return root
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode
	var result [][]int

	queue = append(queue, root)
	result = append(result, []int{root.Val})

	required := 2
	scan := 1

	var arr []int
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			arr = append(arr, node.Left.Val)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			arr = append(arr, node.Right.Val)
			queue = append(queue, node.Right)
		}
		scan++
		if len(arr) != 0 && scan == required {
			result = append(result, arr)
			required = required * 2
			arr = []int{}
		}
	}

	return result
}

func main() {
	// root := []int{3, 9, 20, -1, -1, 15, 7}
	root := []int{1, 2, 3, 4, -1, -1, 5, 1, 1, 1, 1, 1, 1}
	fmt.Println(root)
	rootTree := NewTreeFromArray(root)
	fmt.Println("---print---")
	rootTree.Print()
	fmt.Println("---")
	levelOrder := levelOrder(rootTree)
	fmt.Println(levelOrder)

	// 	1
	// 2 		3
	// 4 -1 -1 5
}

// issue with level by level, not left to right etc
