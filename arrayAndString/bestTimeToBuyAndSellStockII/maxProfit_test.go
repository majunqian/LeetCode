package bestTimeToBuyAndSellStockII

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name string
		prices []int
		expectedProfit int
	}{
		{"",[]int{7,1,5,3,6,4},7},
		{"",[]int{1,2,3,4,5},4},
		{"",[]int{7,6,4,3,1},0},
		{"",[]int{2,1,2,0,1},2},

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxProfit(test.prices)
			if result != test.expectedProfit{
				t.Errorf("%v  expected %d got %d",
					test.prices, test.expectedProfit, result)
			}
		})
	}
}
