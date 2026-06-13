package removellelements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil || (head.Val == val && head.Next == nil) {
		return nil
	}

	dummyNode := &ListNode{Next: head}

	current := dummyNode
	for current != nil {
		nextNode := current.Next
		if nextNode == nil {
			break
		}
		nextVal := nextNode.Val
		if nextVal == val {
			newNext := current.Next.Next
			if nextNode == head {
				head = newNext
			}
			current.Next = newNext
		} else {
			current = current.Next
		}
	}
	return head
}

/*
Notes:

Problem Link: https://leetcode.com/problems/remove-linked-list-elements

Comments:

1. Used a dummy note to account for the head being a node that needs
to be removed.

2. current = current.Next will be done only if the next node is not
removed. This is to account for consecutive removable nodes. If this
is done for every iteration, then consecutive candidates will not be
fully removed, one will remain.

3. The loop goes on till current != nil

*/
