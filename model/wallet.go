package model

type Account struct {
	PrivateKey string
	Address    string
}

type WalletInfo struct {
	WalletAddr  string `json:"walletAddr"`
	BnbBalance  string `json:"bnbBalance"`
	USDCBalance string `json:"usdcBalance"`
	SKSBalance  string `json:"sksBalance"`
}
