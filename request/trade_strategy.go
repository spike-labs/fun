package request

type TradeStrategy struct {
	Price             map[int]string
	Time              map[int]string
	Amount            map[int]string
	Frequency         string
	SlippageTolerance string
	Work              bool
}

type TradeStrategyService struct {
	BuyPrice  string `json:"buy_price" form:"buy_price"`
	SellPrice string `json:"sell_price" form:"sell_price"`
	Time      map[int][]string

	TotalAmount       string `json:"total_amount" form:"total_amount"`
	SigleAmount       string `json:"sigle_amount" form:"sigle_amount"`
	Frequency         string `json:"frequency" form:"frequency"`
	SlippageTolerance string `json:"slippage_tolerance" form:"slippage_tolerance"`
	Work              bool
}
