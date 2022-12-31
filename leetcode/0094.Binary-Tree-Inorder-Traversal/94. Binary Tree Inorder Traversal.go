package leetcode

import (
	"github.com/ru7dy/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var result []int
    inorder(root, &result)
    return result
}

func inorder(root *TreeNode, output *[]int) {
    if root != nil {
        inorder(root.Left, output)
        *output = append(*output, root.Val)
        inorder(root.Right, output)
    }
}

func inorderTraversalInterative(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	current := root 
    
	// stack := []*TreeNode{} 
    stack := make([]*TreeNode, 0)    
	
    for len(stack) > 0 || current != nil {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			last_idx := len(stack) - 1             
			current = stack[last_idx]
			stack = stack[:last_idx]            
			result = append(result, current.Val)
			current = current.Right
		}
	}

	return result
}
