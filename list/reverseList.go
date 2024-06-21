package list

func ReverseList(head *Node) *Node {
	var preNode *Node
	for head != nil {
		//存当前的节点
		currentNode := head
		//切换到下一个节点
		head = head.Next
		//转义关系
		currentNode.Next = preNode
		//修改pre
		preNode = currentNode
	}
	return preNode
}
