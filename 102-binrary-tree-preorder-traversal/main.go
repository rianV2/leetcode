package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func preorderTraversal(node *TreeNode) []int {
	result := []int{}

	if node == nil {
		return []int{}
	}

	result = append(result, node.Val)

	if node.Left != nil {
		result = append(result, preorderTraversal(node.Left)...)
	}
	if node.Right != nil {
		result = append(result, preorderTraversal(node.Right)...)
	}

	return result
}

func main() {
	// [1,2,3,4,5,null,8,null,null,6,7,9]
}
