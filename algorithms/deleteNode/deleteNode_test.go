package deleteNode

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name       string
		listValues []int
		node       int
		expected   []int
	}{
		{"", []int{4, 5, 1, 9}, 1, []int{4, 5, 9}},
		{"", []int{4, 5, 1, 9}, 5, []int{4, 1, 9}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := arrayToList(test.listValues)
			deleteNode(getNode(list, test.node))
			array := listToArray(list)
			if len(array) != len(test.expected) {
				t.Errorf("list %v deleteNode %d got %v expected %v",
					test.listValues, test.node, array, test.expected)
			}
			for i := range array {
				if array[i] != test.expected[i] {
					t.Errorf("list %v deleteNode %d got %v expected %v",
						test.listValues, test.node, array, test.expected)
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

func getNode(head *ListNode, val int) *ListNode {
	for node := head; node != nil; node = node.Next {
		if node.Val == val {
			return node
		}
	}
	return nil
}
