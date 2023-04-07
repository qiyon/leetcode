package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}

	invertTree(root)

	out, _ := json.Marshal(root)
	fmt.Printf("%s", string(out))
}

// leetcode start

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	loopInvert(root)
	return root
}

func loopInvert(root *TreeNode) {
	if nil != root {
		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp

		invertTree(root.Left)
		invertTree(root.Right)
	}
}
