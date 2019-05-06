package ch2

//MaxProfitOnce - calculate profit of buying stock once.
func MaxProfitOnce(prices []int) int {
	min := prices[0]
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		profit := prices[i] - min
		if profit > maxprofit {
			maxprofit = profit
		}
	}
	return maxprofit
}
