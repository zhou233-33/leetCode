package array

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	right := -1
	minlen := len(nums)
	for right < len(nums) {
		if sum < target {
			right++
			if right > len(nums)-1 {
				break
			}
			sum += nums[right]
		} else {
			minlen = Min(right-left+1, minlen)
			sum -= nums[left]
			left++
		}
	}
	return minlen
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
