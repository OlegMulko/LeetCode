package mergetwosortedlists

import "testing"

type testPair struct {
	descriptoin string
	l1          *ListNode
	l2          *ListNode
	average     *ListNode
}

var tests = []testPair{
	{"test1", nil, nil, nil},
	{"test2", GetListNodes(1, 2, 4), nil, GetListNodes(1, 2, 4)},
	{"test3", nil, GetListNodes(1, 2, 4), GetListNodes(1, 2, 4)},
	{"test4", GetListNodes(1, 2, 4), GetListNodes(1, 3, 4), GetListNodes(1, 1, 2, 3, 4, 4)},
}

func TestReverseList(t *testing.T) {
	error := false
	for _, pair := range tests {
		listTest := pair.average
		result := MergeTwoLists(pair.l1, pair.l2)
		for result != nil {
			if listTest.Val != result.Val {
				error = true
				break
			}
			result = result.Next
			listTest = listTest.Next
		}
		if error {
			t.Error(
				"description", pair.descriptoin,
			)
		}
	}
}

func GetListNodes(values ...int) *ListNode {
	var head *ListNode
	var list *ListNode
	var nextList *ListNode

	if len(values) == 0 {
		return nil
	}
	for index, value := range values {
		nextList = &ListNode{}
		nextList.Val = value
		if index == 0 {
			head = nextList
			list = nextList
			continue
		}
		list.Next = nextList
		list = nextList
	}
	return head
}
