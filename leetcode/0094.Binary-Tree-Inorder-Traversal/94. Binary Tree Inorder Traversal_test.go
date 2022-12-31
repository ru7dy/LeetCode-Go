package leetcode

import (
	"fmt"
	"testing"
	"github.com/ru7dy/LeetCode-Go/structures"
)

type question94 struct {
	para94
	ans94
}

// one 代表第一个参数
type para94 struct {
	one []int
}

// one 代表第一个答案
type ans94 struct {
	one []int
}

func Test_Problem94(t *testing.T) {
	qs := []question94{
		{
			para94{[]int{}},
			ans94{[]int{}},
		},
		{
			para94{[]int{1}},
			ans94{[]int{1}},
		},
		{
			para94{[]int{1, structures.NULL, 2, 3}},
			ans94{[]int{1, 2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 94------------------------\n")

	for _, q := range qs {
		_, p := q.ans94, q.para94
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)		
		fmt.Printf("【output】:%v      \n", inorderTraversal(root))
	}	
	fmt.Printf("\n\n\n")
}