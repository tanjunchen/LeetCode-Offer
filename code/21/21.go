package main

// 调整数组顺序使奇数位于偶数前面

// 头尾指针
func exchange(nums []int) []int {
	start, end := 0, len(nums)-1
	for start < end {
		for nums[start]%2 != 0 && start < end {
			start++
		}
		for nums[end]%2 == 0 && start < end {
			end--
		}
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return nums
}

//
func exchange2(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
