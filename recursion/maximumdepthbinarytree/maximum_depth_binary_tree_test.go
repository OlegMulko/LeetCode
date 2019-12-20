package maximumdepthbinarytree

import "testing"

type testPair struct {
	descriptoin string
	value       []int
	average     int
}

var tests = []testPair{
	{"MaxDepth = 0", []int{0}, 0},
	{"MaxDepth = 1", []int{20, 0, 0}, 1},
	{"MaxDepth = 2 left", []int{20, 11, 0, 0, 0}, 2},
	{"MaxDepth = 2 right", []int{20, 0, 11, 0, 0}, 2},
	{"MaxDepth = 3 right", []int{3, 9, 0, 0, 20, 15, 0, 0, 7, 0, 0}, 3},
	{"MaxDepth = 3 right", []int{3, 20, 15, 0, 0, 7, 0, 0, 9, 0, 0}, 3},
}

func TestMaxDepth(t *testing.T) {
	var tree *TreeNode

	error := false
	for _, pair := range tests {
		index := 0
		tree = generateBinaryTree(pair.value, nil, &index)
		result := MaxDepth(tree)
		if result != pair.average {
			error = true
		}
		if error {
			t.Error(
				"description", pair.descriptoin,
			)
		}
	}
}

func generateBinaryTree(binaryTree []int, tree *TreeNode, index *int) *TreeNode {

	lenBT := len(binaryTree)

	if *index >= lenBT {
		return nil
	}
	if tree == nil {
		if tree = getNode(binaryTree[*index]); tree == nil {
			return nil
		}
	}
	(*index)++
	if *index >= lenBT {
		return nil
	}
	tree.Left = generateBinaryTree(binaryTree, tree.Left, index)
	(*index)++
	if *index >= lenBT {
		return nil
	}
	tree.Right = generateBinaryTree(binaryTree, tree.Right, index)
	return tree
}

func getNode(value int) *TreeNode {
	if value == 0 {
		return nil
	}
	return &TreeNode{
		Val: value,
	}
}
