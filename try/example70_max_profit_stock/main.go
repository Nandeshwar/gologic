package main

// stocks - is price for each day.whey you buy 1 stock  and sell it, you will  get max profit
// 10, 2, 3, 10, 5, 7
// consider 1st item minSellPrice
// Iterate and find profit and also keep maintaining lowest price so that it will be easy find profit
func findMaxProfit(stocks []int) int {
	minSellPrice := stocks[0]
	maxProfit := 0
	
	profit := 0
	
	for _, v := range stocks {
		if v < minSellPrice {
			minSellPrice = v
		}
		
		profit = v - minSellPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}