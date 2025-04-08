package main

func BreadthFirstTraversal(head TreeNode[int]) []int {
	queue := []TreeNode[int]{head}
	path := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		path = append(path, node.value)

		if node.left != nil {
			queue = append(queue, *node.left)
		}

		if node.right != nil {
			queue = append(queue, *node.right)
		}
	}

	return path
}

func BreadthFirstSearch(head TreeNode[int], target int) bool {
	queue := []TreeNode[int]{head}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.value == target {
			return true
		}

		if node.left != nil {
			queue = append(queue, *node.left)
		}

		if node.right != nil {
			queue = append(queue, *node.right)
		}
	}

	return false
}
