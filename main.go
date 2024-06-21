package main

import "zhouhzLearn/list"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	linkList := list.ConvertArrayToLinkedList(nums)
	res := list.RemoveNthFromEnd(linkList, 3)
	list.PrintLinkedList(res)
}
