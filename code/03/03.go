package main

import "sort"

/**
offer-03 数组中的重复数字
*/

/**
解法一 哈希
*/
func findRepeatNumber(nums []int) int {
	if nums == nil {
		return -1
	}
	repeatedMap := make(map[int]int)
	for _, i := range nums {
		if _, ok := repeatedMap[i]; ok {
			return i
		}
		repeatedMap[i] = 1
	}
	return -1
}

/**
解法二 利用数字在 0~n-1 之间的规律
*/
func findRepeatNumber2(nums []int) int {
	if nums == nil {
		return -1
	}
	length := len(nums)
	for i := 0; i < length; i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return -1
}

/**
解法三 先排序, 后查找
*/
func findRepeatNumber3(nums []int) int {
	if nums == nil {
		return -1
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
