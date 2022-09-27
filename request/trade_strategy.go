package request

type TradeStrategy struct {
	Price             map[int]string
	ExecTime          int64
	Amount            map[int]string
	Frequency         int64
	SlippageTolerance int64
	Work              bool

	Closing chan struct{}
}

type TradeStrategyService struct {
	// Purchase price, low price
	BuyPrice string `json:"buy_price" form:"buy_price"`
	// Sell price ,high price
	SellPrice string `json:"sell_price" form:"sell_price"`
	// Strategy Runtime Unit Minutes
	ExecTime int64 `json:"exec_time" form:"exec_time"`
	// The total amount spent on this strategy, unit U
	TotalAmount string `json:"total_amount" form:"total_amount"`
	// The single amount spent on this strategy, unit U
	SingleAmount string `json:"single_amount" form:"single_amount"`
	// The frequency of execution of the policy, that is, the interval between two purchases
	Frequency int64 `json:"frequency" form:"frequency"`
	// The sliding point of this transaction, set the minimum amount obtained by this flash exchange through the subparameter.Unit: kilobit
	SlippageTolerance int64 `json:"slippage_tolerance" form:"slippage_tolerance"`
	Work              bool
}

type DelTradeStrategyService struct {
	Uuid string `json:"uuid" form:"uuid"`
}
