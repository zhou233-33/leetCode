package array

func FindTargetIndex(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = left + 1
		}
		if nums[mid] > target {
			right = right - 1
		}
	}
	return -1
}
func FindMaxMin(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	result := -1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			result = mid
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] == target {
			return mid
		}
	}
	return result
}
func FindRightIndex(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	result := -1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] <= target {
			result = mid
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return result
}
func FindLeftIndex(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	result := -1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] >= target {
			result = mid
			right = mid - 1
		}
	}
	return result
}
