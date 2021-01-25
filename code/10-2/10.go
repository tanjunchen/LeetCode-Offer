package main

// 青蛙跳台阶问题
func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	var res = make([]int, n)
	res[0] = 1
	res[1] = 2
	for i := 2; i < n; i++ {
		res[i] = res[i-1]%1000000007 + res[i-2]%1000000007
	}
	return res[n-1] % 1000000007
}

func numWays2(n int) int {
	f1, f2 := 1, 1
	for i := 1; i <= n; i++ {
		f1, f2 = f2, (f1+f2)%1000000007
	}
	return f1
}
