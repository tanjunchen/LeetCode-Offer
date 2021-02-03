package main

// 求1+2+…+n
// 求 1+2+...+n , 要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句(A?B:C)
func sumNums(n int) int {
	return n * (n + 1) / 2
}

func sumNums2(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumNums(n-1)
}

func sumNums3(n int) int {
	res := 0
	var sum func(int) bool
	sum = func(n int) bool {
		res += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return res
}
