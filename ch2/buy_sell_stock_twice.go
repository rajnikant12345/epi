package ch2

//MaxProfitTwice calculate max profit for buying and selling stock twice
func MaxProfitTwice(prices []int) int {
	profitArray1 := make([]int, len(prices))
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
		profitArray1[i] = maxprofit
	}

	max := prices[len(prices)-1]
	maxprofit = 0
	maxprofitout := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		profit := max - prices[i]
		if profit > maxprofit {
			maxprofit = profit
		}
		if i != 0 {
			tmp := maxprofit + profitArray1[i-1]
			if tmp > maxprofitout {
				maxprofitout = tmp
			}
		} else {
			tmp := maxprofit
			if tmp > maxprofitout {
				maxprofitout = tmp
			}
		}

	}
	return maxprofitout
}
