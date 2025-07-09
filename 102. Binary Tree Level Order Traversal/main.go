package main

import (
	"fmt"
)

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

/*
History
 1. add to array current val only, not his left and right directly
    they will be added at next queue scan
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode
	var result [][]int

	queue = append(queue, root)
	// result = append(result, []int{root.Val}) //removed since, we scan it from queue

	// required := 2
	// scan := 1 //had scan the root

	for len(queue) > 0 { //O(n)
		currentQSize := len(queue)
		var arr []int

		/* scan current size queue first
		while scanning current queue size, it will only add new queue
		without executing it directly
		*/
		for i := 0; i < currentQSize; i++ { //O(n) constant factor cause depend on populated data (qeueu)
			node := queue[0]
			queue = queue[1:]
			arr = append(arr, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, arr)
	}

	/*
		Time Complexity: O(n)
		Space Complexity: O(w)
			- queue added, and removed
			- this will depend on the longest size of a level
	*/

	return result
}

func main() {
	//incomplete binary tree, left side only
	root := []int{1, 2, -1, 3, -1, 4, -1, 5}

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
