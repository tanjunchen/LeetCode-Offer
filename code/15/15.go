package main

// 二进制数中 1 的个数
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = (num - 1) & num
	}
	return count
}

func hammingWeight2(num uint32) int {
	count := 0
	var flag uint32 = 1
	for flag != 0 {
		if num != 0&flag {
			count++
		}
		flag = flag << 1
	}
	return count
}
