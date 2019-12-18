package reverselinkedlist

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList ...
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := head
	nextList := head.Next
	head = ReverseList(nextList)
	if nextList == nil {
		return list
	}
	nextList.Next = list
	list.Next = nil
	return head
}
