package model

type Account struct {
	PrivateKey string
	Address    string
}

type WalletInfo struct {
	WalletAddr             string `json:"walletAddr"`
	BnbBalance             string `json:"bnbBalance"`
	USDCBalance            string `json:"usdcBalance"`
	SKSBalance             string `json:"sksBalance"`
	SKSEarningsToday       string `json:"sksEarningsToday"`
	USDCEarningsToday      string `json:"usdcEarningsToday"`
	SKSCumulativeEarnings  string `json:"sksCumulativeEarnings"`
	USDCCumulativeEarnings string `json:"usdcCumulativeEarnings"`
}
