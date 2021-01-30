package main

import "math"

// 剪绳子
func cuttingRope(n int) int {
	if n == 2 || n == 1 {
		return 1
	}
	if n == 3 {
		return 2
	}
	pow := n / 3
	value := n % 3
	var res float64 = 1
	switch value {
	case 1:
		res *= math.Pow(3, float64(pow-1))
		res *= 4
	case 2:
		res *= math.Pow(3, float64(pow))
		res *= 2
	default:
		res *= math.Pow(3, float64(pow))
	}
	return int(res)
}
