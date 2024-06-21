package list

func RemoveNthFromEnd(head *Node, n int) *Node {
	fastIndex, slowIndex := head, head
	slowPre := &Node{
		Data: 0,
		Next: head,
	}
	result := slowPre
	step := 0
	for fastIndex != nil {
		step++
		if step > n {
			slowPre = slowIndex
			slowIndex = slowIndex.Next
		}
		fastIndex = fastIndex.Next
	}
	slowPre.Next = slowIndex.Next
	return result.Next
}
