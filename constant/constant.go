package constant

const (
	ConfigFile          = "config.toml"
	ConfigEnv           = "ENV_CONFIG"
	PROTOCOL            = "https"
	DOMAIN              = "deep-index.moralis.io"
	MORALIS_API_VERSION = "api/v2"
	MORALIS_API         = PROTOCOL + "://" + DOMAIN + "/" + MORALIS_API_VERSION + "/"
	NewBlockTopic       = "newBlockTopic"
	BlockConfirmHeight  = 15
	EmptyAddress        = "0x0000000000000000000000000000000000000000"
)

const (
	BUY_PRICE = iota + 1
	SELL_PRICE
	TOTAL_AMOUNT
	SINGLE_AMOUNT
)
