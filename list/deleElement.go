package list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// 将数组转换为链表的函数
func ConvertArrayToLinkedList(array []int) *Node {
	if len(array) == 0 {
		return nil
	}

	// 初始化链表头节点
	head := &Node{Data: array[0]}
	current := head

	// 遍历数组，创建链表节点并链接
	for i := 1; i < len(array); i++ {
		newNode := &Node{Data: array[i]}
		current.Next = newNode
		current = newNode
	}
	return head
}

func DeleteLinkedListItem(head *Node, target int) *Node {
	preNode := &Node{
		Data: 0,
		Next: head,
	}
	result := preNode
	StepNode := head
	for StepNode != nil {
		if StepNode.Data == target {
			preNode.Next = StepNode.Next
		} else {
			preNode = StepNode
		}
		StepNode = StepNode.Next
	}
	return result.Next
}
func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
