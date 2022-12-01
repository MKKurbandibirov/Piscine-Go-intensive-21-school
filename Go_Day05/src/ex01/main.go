package main

import (
	"day05/domain"
	"fmt"
)

func levelOrderTraversal(root *domain.TreeNode) [][]bool {
	if root == nil {
		return nil
	}
	queue := make([]*domain.TreeNode, 0)
	queue = append(queue, root)
	res := make([][]bool, 0)
	level := 0
	for len(queue) > 0 {
		i := 0
		l := len(queue)
		res = append(res, make([]bool, 0))
		for ; i < l; i++ {
			node := queue[i]
			res[level] = append(res[level], node.HasToy)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[i:]
		level++
	}
	return res
}

func unrollGarland(tree *domain.BinaryTree) []bool {
	traversal := levelOrderTraversal(tree.Root)
	result := make([]bool, 0)
	for i := 1; i <= len(traversal); i++ {
		if i%2 == 0 {
			for j := 0; j < len(traversal[i-1]); j++ {
				result = append(result, traversal[i-1][j])
			}
		} else {
			for j := len(traversal[i-1]) - 1; j >= 0; j-- {
				result = append(result, traversal[i-1][j])
			}
		}
	}
	return result
}

func main() {
	tree := domain.NewBinaryTree(true)

	tree.Root.Left = domain.NewTreeNode(true)
	tree.Root.Left.Right = domain.NewTreeNode(false)
	tree.Root.Left.Left = domain.NewTreeNode(true)

	tree.Root.Right = domain.NewTreeNode(false)
	tree.Root.Right.Right = domain.NewTreeNode(true)
	tree.Root.Right.Left = domain.NewTreeNode(true)

	fmt.Println(unrollGarland(tree))
}
