package addTwoNumbers

import (
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listToNum(l *ListNode) *big.Int {
	var (
		num big.Int
		sum string
	)
	p := l
	for {
		sum = strconv.Itoa(p.Val) + sum
		if p.Next != nil {
			p = p.Next
		} else {
			num.SetString(sum, 10)
			return &num
		}
	}
}
func numToList(n *big.Int) *ListNode {
	s := n.String()
	lists := make([]ListNode, len(s))

	for i, v := range []byte(s) {
		lists[len(lists)-1-i].Val, _ = strconv.Atoi(string(v))
		if i != 0 {
			lists[len(lists)-1-i].Next = &lists[len(lists)-i]
		}
	}
	return &(lists[0])
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c big.Int
	c.Add(listToNum(l2), listToNum(l1))
	return numToList(&c)
}
