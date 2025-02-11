package _160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// ways 01 use linked list
	//if headA == nil || headB == nil {
	//	return nil
	//}
	//dummy1, dummy2 := headA, headB
	//for dummy1 != dummy2 {
	//	if dummy1 != nil {
	//		dummy1 = dummy1.Next
	//	} else {
	//		dummy1 = headB
	//	}
	//	if dummy2 != nil {
	//		dummy2 = dummy2.Next
	//	} else {
	//		dummy2 = headA
	//	}
	//}
	//return dummy1

	// ways 02 use hash table

	hashMap := make(map[*ListNode]struct{})

	for node := headA; node != nil; node = node.Next {
		hashMap[node] = struct{}{}
	}

	for node := headB; node != nil; node = node.Next {
		if _, ok := hashMap[node]; ok {
			return node
		}
	}
	return nil
}

func Run(headA, headB *ListNode) *ListNode {
	return getIntersectionNode(headA, headB)
}
