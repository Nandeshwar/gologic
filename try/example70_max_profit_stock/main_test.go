package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfitInStockPurchaseAndSale(t *testing.T) {
	Convey("Test Max profit in stocks", t, func() {
		Convey("success testing max profit", func() {
			data := []struct {
				stockPriceList []int
				maxProfit      int
			}{
				{stockPriceList: []int{10, 2, 3, 10, 5, 7}, maxProfit: 8},
			}

			for _, d := range data {
				So(findMaxProfit(d.stockPriceList), ShouldEqual, d.maxProfit)
			}
		})
	})
}
