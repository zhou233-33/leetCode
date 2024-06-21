package array

import (
	"math"
	"sort"
)

func SortedSquares(nums []int) []int {
	for i, val := range nums {
		nums[i] *= val
	}
	sort.Ints(nums)
	return nums
}

func SortedSquaresByDoublePtr(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, 0, len(nums))
	for left <= right {
		if math.Abs(float64(nums[left])) <= math.Abs(float64(nums[right])) {
			res = append([]int{nums[right] * nums[right]}, res...)
			right--
		} else {
			res = append([]int{nums[left] * nums[left]}, res...)
			left++
		}
	}
	return res
}
