package list

func GetIntersectionNode(headA, headB *Node) *Node {
	ptr1, ptr2 := headA, headB
	for ptr2 != ptr1 {
		if ptr1 != nil {
			ptr1 = ptr1.Next
		} else {
			ptr1 = headB
		}
		if ptr2 != nil {
			ptr2 = ptr2.Next
		} else {
			ptr2 = headA
		}
	}
	//如果不相交，可以看作相交的点是nil
	return ptr1
}
