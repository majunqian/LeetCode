package reverseList

import "testing"

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		lists    []int
		expected []int
	}{
		{"", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"", []int{}, []int{}},
		{"", []int{2}, []int{2}},
		{"", []int{1, 2}, []int{2, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			listNodes := arrayToList(test.lists)
			reversedList := reverseList(listNodes)
			result := listToArray(reversedList)
			if len(test.expected) != len(result) {
				t.Errorf("list:%v  got:%v expected:%v", test.lists, result, test.expected)
				return
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("list:%v got:%v expected:%v", test.lists, result, test.expected)
					break
				}
			}
		})
	}
}

func arrayToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
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
