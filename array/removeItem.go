package array

func RemoveItem(nums []int, target int) []int {
	slow := 0
	fast := 0
	count := 0
	for fast < len(nums) {
		//如果slow是target
		if nums[fast] == target {
			count++
		}
		if nums[slow] == target && nums[fast] != target {
			nums[fast], nums[slow] = nums[slow], nums[fast]
		}
		fast++
		if nums[slow] == target {
			continue
		} else {
			slow++
		}
	}
	return nums[:len(nums)-count]
}
