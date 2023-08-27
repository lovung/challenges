package contest

// https://leetcode.com/problems/account-balance-after-rounded-purchase/
func accountBalanceAfterPurchase(purchaseAmount int) int {
	if purchaseAmount%10 < 5 {
		purchaseAmount /= 10
	} else {
		purchaseAmount /= 10
		purchaseAmount++
	}
	return 100 - purchaseAmount*10
}
