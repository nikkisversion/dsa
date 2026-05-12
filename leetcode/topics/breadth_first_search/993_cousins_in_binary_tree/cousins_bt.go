package cousinsinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {

	if root == nil {
		return false
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {

		l := len(queue)
		foundX := false
		foundY := false

		for i := 0; i < l; i++ {

			current := queue[0]
			queue[0] = nil
			queue = queue[1:]

			left := current.Left
			right := current.Right

			if left != nil && right != nil {
				if (left.Val == x || left.Val == y) && (right.Val == x || right.Val == y) {
					return false
				}
			}

			if left != nil {
				queue = append(queue, left)
			}

			if right != nil {
				queue = append(queue, right)
			}

			if current.Val == x {
				foundX = true
			}

			if current.Val == y {
				foundY = true
			}

			if foundX && foundY {
				return true
			}

		}

	}
	return false
}

/*
Notes

Problem Link:
https://leetcode.com/problems/cousins-in-binary-tree

Comments:

1. We can do this via DFS or BFS.

2. Here with BFS, we are first checking if the current
node has two children with values equal to x and y. If
so, then the answer would be false as by the problem
statement, cousin nodes cannot have the same parent.
This check ensures that w do this for each node.

3. Since in each iteration of the queue (the loop with
i as iterator), we are checking all nodes at one level.
Here now we can do the second check of values x and y.
As we have already eliminated the possibility of same
parent in above check, we can now safely assume that
if we find x and y both at this level, then the code
will return true.

*/
