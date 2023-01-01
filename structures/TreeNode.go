package structures

// import (
// 	"fmt"
// 	"strconv"
// )

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Ins2Treenode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil 
	}
	
	root := &TreeNode{Val: ints[0]}
	queue := make([]*TreeNode, 1, n*2) //len 1, cap 2n
	queue[0] = root 

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++ 

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root 
}

// T2s converts *TreeNode to *[]int
// preorder? 
func T2s(head *TreeNode, array *[]int) {
	*array = append(*array, head.Val)
	if head.Left != nil {
		T2s(head.Left, array)
	}
	if head.Right != nil {
		T2s(head.Right, array)
	}
}
