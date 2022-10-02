package main

import "math"

// 4.4 Check Balanced: Implement a function to check if a binary tree is balanced.
// For the purposes of this question, a balanced tree is defined to be a tree such that
// the heights of the two subtrees of any node never differ by more than one

func isBalanced(root *NodeInt) bool {
	var (
		leftCount  int
		rightCount int
	)
	leftCount = countLevel(root.LeftNode, 1)
	rightCount = countLevel(root.RightNode, 1)

	diff := leftCount - rightCount
	diff = int(math.Abs(float64(diff)))
	return diff <= 1
}

func countLevel(node *NodeInt, level int) int {
	if node.LeftNode != nil && node.RightNode != nil {
		level += 1
	}
	if node.LeftNode == nil && node.RightNode != nil {
		level += 1
	}
	if node.LeftNode != nil && node.RightNode == nil {
		level += 1
	}

	if node.LeftNode != nil {
		level = countLevel(node.LeftNode, level)
	}
	if node.RightNode != nil {
		level = countLevel(node.RightNode, level)
	}

	return level
}

//4.5 Validate BST: Implement a function to check if a binary tree is a binary search tree

func isBST(root *NodeInt) (bool, int) {
	if root == nil {
		return true, -1
	}

	isBstLeft, leftValue := isBST(root.LeftNode)
	if leftValue == -1 {
		return true, leftValue
	}
	if !isBstLeft && leftValue != -1 {
		return false, leftValue
	}
	value := root.Value
	isBstRight, rightValue := isBST(root.RightNode)
	if rightValue == -1 {
		return true, rightValue
	}
	if !isBstRight && rightValue != -1 {
		return false, rightValue
	}

	if leftValue > value || rightValue <= value {
		return false, 0
	}

	return isBstLeft && isBstRight, value
}
