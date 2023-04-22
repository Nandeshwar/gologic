package main

import (
	"fmt"
)

/*
Lets understand with an example:

length   | 1   2   3   4   5   6   7   8
--------------------------------------------
price    | 1   5   8   9  10  17  17  20
Given a rod of length of n (here n=8) and its prices if it is sold in that many piece.
That is:

 If we sell 1 piece its price will be 1
 if we sell 2 piece its price will be 5
 Similarly:

	Pieces       Cost
	---------------------
	 3 piece =    8
	 4 piece =    9
	 5 piece =   10
	 6 piece =   17
	 7 piece =   17
	 8 piece =   20
Now we have to find the way to cut the rod in pieces such that the profit obtained when sold becomes maximum.

Approach:
Now in the above example :
8 pieces = n
and its cost is 20. That is if the whole rod is sold without any cuts profit will be 20.

But is it the maximum profit , lets check:

Now :
Let's sell the rod in

pieces    cost
----------------
6 pieces = 17
2 pieces = 5
Total Profit  : 17+5 = 22
*/

func main() {
	// length: 1  2  3  4   5  6   7   8
	a := []int{1, 5, 8, 9, 10, 17, 17, 20} // expectation - 22
	rodLen := 8

	fmt.Println(rodCuttingMaxProfit(a, len(a)-1, rodLen))
}

func rodCuttingMaxProfit(a []int, ind, maxRodLen int) int {

	if ind == 0 {
		return maxRodLen * a[0]
	}

	notPick := 0 + rodCuttingMaxProfit(a, ind-1, maxRodLen)

	rodLen := ind + 1

	pick := -100
	if rodLen <= maxRodLen {
		pick = a[ind] + rodCuttingMaxProfit(a, ind-1, maxRodLen-rodLen)
	}

	return max(pick, notPick)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
