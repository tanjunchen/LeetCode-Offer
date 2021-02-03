package main

// 剑指 Offer 63. 股票的最大利润
// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var min, max int
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			min = prices[0]
			continue
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}



