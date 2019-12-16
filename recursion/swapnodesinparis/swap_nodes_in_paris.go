package swapnodesinparis

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// SwapPairs ...
func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	firstList := head
	secondList := head.Next
	if secondList == nil {
		return head
	}
	val := firstList.Val
	firstList.Val = secondList.Val
	secondList.Val = val
	SwapPairs(secondList.Next)
	return head
}
