package removeNthNodeFromEndOfList

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		n        int
		expected []int
	}{
		{"", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"", []int{1, 2, 3}, 2, []int{1, 3}},
		{"", []int{1, 2}, 2, []int{2}},
		{"", []int{1, 2}, 1, []int{1}},
		{"", []int{1}, 1, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := arrayToList(test.list)
			resultList := removeNthFromEnd(list, test.n)
			resultArr := listToArray(resultList)
			if len(test.expected) != len(resultArr) {
				t.Errorf("list:%v n=%d got:%v expected:%v", test.list, test.n, resultArr, test.expected)
				return
			}
			for i := range resultArr {
				if resultArr[i] != test.expected[i] {
					t.Errorf("list:%v n=%d got:%v expected:%v", test.list, test.n, resultArr, test.expected)
					break
				}
			}
		})
	}
}

func arrayToList(arr []int) *ListNode {
	list := make([]ListNode, len(arr))
	for i := range arr {
		if len(arr)-1 == i {
			list[i].Next = nil
		} else {
			list[i].Next = &list[i+1]
		}
		list[i].Val = arr[i]
	}
	return &list[0]
}
func listToArray(head *ListNode) []int {
	var arr []int
	for node := head; node != nil; node = node.Next {
		arr = append(arr, node.Val)
	}
	return arr
}
