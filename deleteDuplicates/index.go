package deleteduplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// way1
	//for node := head; node != nil; node = node.Next {
	//	for node.Next != nil && node.Val == node.Next.Val {
	//		node.Next = node.Next.Next
	//	}
	//}
	//return head
	if head == nil {
		return nil
	}

	//way2
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func Run(head *ListNode) *ListNode {
	return deleteDuplicates(head)
}
