package domain

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func NewTreeNode(value bool) *TreeNode {
	return &TreeNode{
		HasToy: value,
		Left:   nil,
		Right:  nil,
	}
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree(rootValue bool) *BinaryTree {
	return &BinaryTree{
		Root: NewTreeNode(rootValue),
	}
}
