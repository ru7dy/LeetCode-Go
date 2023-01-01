package leetcode

import (
	"github.com/ru7dy/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue := []*TreeNode{root}
    res := make([][]int, 0)
    for len(queue) >0 {
        l := len(queue)
        tmp := make([]int, 0, l)
        // add current level nodes into res and push next level nodes into queue
        for i := 0; i < l; i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
            tmp = append(tmp, queue[i].Val) 
        }
        // pop up processed nodes on the current level
        queue = queue[l:]
        res = append(res, tmp)
    }
    return res 
}

// DFS
func levelOrderDFS(root *TreeNode) [][]int {
    var res [][]int
    var dfsLevel func(node *TreeNode, level int)
    // func literals are closures 
    dfsLevel = func(node *TreeNode, level int) {
        if node == nil {
            return 
        }
        if len(res) == level {
            // first time to the new level, initialize the inner array with the leftmost val
            res = append(res, []int{node.Val})
        } else {
            // trick: lenth now is +1 to the level when it's not first time encoutered, append the to pre-existing inner array
            res[level] = append(res[level], node.Val)
        }
        dfsLevel(node.Left, level+1)
        dfsLevel(node.Right, level+1)
    }
    dfsLevel(root, 0)
    return res 
}