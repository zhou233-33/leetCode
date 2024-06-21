package list

func SwapPairs(head *Node) *Node {
	leftPre := &Node{
		Data: 0,
		Next: head,
	}
	result := leftPre
	for leftPre != nil && leftPre.Next != nil && leftPre.Next.Next != nil {
		left := leftPre.Next
		right := leftPre.Next.Next
		left.Next = right.Next
		leftPre.Next = right
		right.Next = left
		leftPre = left
	}
	return result.Next
}
