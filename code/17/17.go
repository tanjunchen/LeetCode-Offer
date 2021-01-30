package main

// 打印从 1 到最大的 n 位数
func printNumbers(n int) []int {
	if n < 0 {
		return []int{}
	}
	maxValue := 1
	for i := 0; i < n; i++ {
		maxValue *= 10
	}
	res := make([]int, maxValue-1)
	for i := 0; i < maxValue-1; i++ {
		res[i] = i + 1
	}
	return res
}
