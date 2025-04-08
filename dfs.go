package main

type TreeNode[T any] struct {
	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func PreOrderTraversal(curr *TreeNode[int], path []int) {
	if curr == nil {
		return
	}
	path = append(path, curr.value)
	PreOrderTraversal(curr.left, path)
	PreOrderTraversal(curr.right, path)
}

func CompareTrees(first *TreeNode[int], second *TreeNode[int]) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil || second == nil {
		return false
	}

	if first.value != second.value {
		return false
	}

	return CompareTrees(first.left, second.left) && CompareTrees(first.right, second.right)
}

func BinaryTreeSearch(curr *TreeNode[int], target int) bool {
	if curr == nil {
		return false
	}

	if curr.value == target {
		return true
	}

	if curr.value < target {
		return BinaryTreeSearch(curr.right, target)
	}

	return BinaryTreeSearch(curr.left, target)
}
