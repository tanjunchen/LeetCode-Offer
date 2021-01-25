package main

import "fmt"

// 旋转数组的最小数字
// o(n)
func minArray(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return -1
	}
	min := numbers[0]
	for i := 1; i < length; i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return min
}

func minArray(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return -1
	}
	l, r := 0, length-1
	for l < r {
		mid := l + (r-l)>>2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}
	return numbers[l]
}

func main() {
	res := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(res))
}
