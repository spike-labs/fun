package request

type BuyNFTService struct {
	// unit hour
	Time int64 `json:"time" form:"time"`
	// leftNFTAmount
	BuyAmount int64 `json:"buy_amount" form:"buy_amount"`
}

type BuyNFTStrategy struct {
	Time       int64
	BuyAmount  int64
	LeftAmount int64
	Closing    chan struct{}
}

type DelBuyStrategyService struct {
	Uuid string `json:"uuid" form:"uuid" binding:"required"`
}
