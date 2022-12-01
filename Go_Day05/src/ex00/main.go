package main

import "day05/domain"

func recursiveCheck(node *domain.TreeNode, currCount int) int {
	if node == nil {
		return currCount
	}

	if node.HasToy == true {
		currCount++
	}
	currCount = recursiveCheck(node.Left, currCount)
	currCount = recursiveCheck(node.Right, currCount)
	return currCount
}

func areToysBalanced(tree *domain.BinaryTree) bool {
	leftCount := recursiveCheck(tree.Root.Left, 0)
	rightCount := recursiveCheck(tree.Root.Right, 0)

	return leftCount == rightCount
}

func main() {
	tree := domain.NewBinaryTree(true)

	tree.Root.Left = domain.NewTreeNode(true)
	tree.Root.Left.Right = domain.NewTreeNode(false)
	tree.Root.Left.Left = domain.NewTreeNode(true)

	tree.Root.Right = domain.NewTreeNode(true)
	tree.Root.Right.Right = domain.NewTreeNode(true)
	tree.Root.Right.Left = domain.NewTreeNode(true)

	areToysBalanced(tree)
}
