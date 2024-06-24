package hash

import "sort"

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, val := range nums {
		if index, ok := m[target-val]; ok {
			return []int{index, i}
		}
		m[val] = i
	}
	return nil
}

// ThreeSum 找到三数之和是target的数组，可以考虑如何将三数问题转为两数问题
// 可以遍历一个数为val,另外两个数之和为target-val,问题就变成了两数问题
// 还有一个去重的问题
func ThreeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		//第一位数去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoTarget := target - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == twoTarget {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				//这个循环只停留在了最后一个相等的元素
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for right > left && nums[left] == nums[left+1] {
					left++
				}
				right--
				left++
			}
			if nums[left]+nums[right] > twoTarget {
				right--
			}
			if nums[left]+nums[right] < twoTarget {
				left++
			}
		}
	}
	return result
}

// fourSumCount 还是将问题转换为两数的问题
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	nums12Map := map[int]int{}
	result := 0
	for _, i1 := range nums1 {
		for _, i2 := range nums2 {
			nums12Map[i1+i2] = nums12Map[i1+i2] + 1
		}
	}
	for _, i3 := range nums3 {
		for _, i4 := range nums4 {
			sum := i3 + i4
			result += nums12Map[0-sum]
		}
	}
	return result
}

// FourSum 四数之和与三数之和的逻辑是相同的
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > 0 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
					continue
				}
				if sum > target {
					right--
				}
				if sum < target {
					left++
				}
			}
		}
	}
	return result
}
