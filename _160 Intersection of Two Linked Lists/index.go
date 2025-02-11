package _160_Intersection_of_Two_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	//dumb,
	return headB
}

func Run(headA, headB *ListNode) *ListNode {
	return getIntersectionNode(headA, headB)
}
