package response

type TradeStrategyResponse struct {
	Uuid string `json:"uuid" form:"uuid"`
	// Purchase price, low price
	BuyPrice string `json:"buy_price" form:"buy_price"`
	// Sell price ,high price
	SellPrice string `json:"sell_price" form:"sell_price"`
	// Strategy Runtime .Unit: Minutes
	ExecTime int64 `json:"exec_time" form:"exec_time"`
	// The total amount spent on this strategy, unit U
	TotalAmount string `json:"total_amount" form:"total_amount"`
	// The single amount spent on this strategy, unit U
	SingleAmount string `json:"single_amount" form:"single_amount"`
	// The frequency of execution of the policy, that is, the interval between two purchases .unit :Second
	Frequency int64 `json:"frequency" form:"frequency"`
	// The sliding point of this transaction, set the minimum amount obtained by this flash exchange through the subparameter.Unit: kilobit
	SlippageTolerance int64 `json:"slippage_tolerance" form:"slippage_tolerance"`
}
