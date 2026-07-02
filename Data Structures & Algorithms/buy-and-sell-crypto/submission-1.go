

func maxProfit(prices []int) int {
	profit := 0
	minPurchase := prices[0]
	for i := 1; i < len(prices); i++ {

		if prices[i] < minPurchase {
			minPurchase = prices[i]
		} else {
			profit = max(prices[i]-minPurchase, profit)
		}
	}
	return profit
}