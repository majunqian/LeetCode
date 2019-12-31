package removeNthNodeFromEndOfList

//Given a linked list, remove the n-th node from the end of list and return its head.
//
//Example:
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//
//Given n will always be valid.
//
//Follow up:
//
//Could you do this in one pass?

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	if len(nodes) == 0 || n < 1 {
		return head
	}
	nTh := len(nodes) - n
	if n == 1 {
		if nTh == 0 {
			nodes[0] = nil
		} else {
			nodes[nTh-1].Next = nil
		}
	} else {
		nodes[nTh].Val = nodes[nTh].Next.Val
		nodes[nTh].Next = nodes[nTh].Next.Next
	}

	return nodes[0]
}
