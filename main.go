package main

import (
	"fmt"
	"zhouhzLearn/hash"
)

func main() {
	nums1 := []int{1, 0, -1, 0, -2, 2}
	res := hash.FourSum(nums1, 0)
	fmt.Println(res)
}
