package binarytreelevelordertraversalbfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := [][]int{}
	level := 0

	for true {

		if len(queue) == 0 {
			break
		}

		l := len(queue)
		levelNodes := []int{}
		for i := 0; i < l; i++ {

			current := queue[0]
			queue[0] = nil
			queue = queue[1:]

			if current == nil {
				continue
			}

			levelNodes = append(levelNodes, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		res = append(res, levelNodes)
		level++
	}

	return res

}

/*
Notes

Problem Link:
https://leetcode.com/problems/binary-tree-level-order-traversal

Comments:

1. The number of elements in the queue is the number of elements at that level.

2. We are setting queue[0] = nil to free up memory to be collected by the GC later.

3. In the loop with iterator as i, we need to calculate the length of the queue
beforehand and use that as the limit for the iterator. This is because when later
we append to the queue, its length will keep increasing and this loop will keep
going on.

*/
