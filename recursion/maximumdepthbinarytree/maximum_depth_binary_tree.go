package maximumdepthbinarytree

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxDepth ...
func MaxDepth(root *TreeNode) int {
	var result int

	if root == nil {
		return 0
	}
	leftResult := MaxDepth(root.Left)
	RightResult := MaxDepth(root.Right)
	if leftResult >= RightResult {
		result = leftResult
	} else {
		result = RightResult
	}
	return 1 + result
}
