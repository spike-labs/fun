package request

type BuyNFTService struct {
	Time      int64 `json:"time" form:"time"`
	BuyAmount int64 `json:"buy_amount" form:"buy_amount"`
}

type BuyNFTStrategy struct {
	Time      int64
	BuyAmount int64
	Closing   chan struct{}
}
