package constant

const (
	ConfigFile          = "config.toml"
	ConfigEnv           = "ENV_CONFIG"
	PROTOCOL            = "https"
	DOMAIN              = "deep-index.moralis.io"
	MORALIS_API_VERSION = "api/v2"
	MORALIS_API         = PROTOCOL + "://" + DOMAIN + "/" + MORALIS_API_VERSION + "/"
	BlockConfirmHeight  = 15
	EmptyAddress        = "0x0000000000000000000000000000000000000000"

	LoginAddr       = "/api/auth/login"
	SendMailAddr    = "/api/user/send/band/email"
	BindAccountAddr = "/api/user/band/address"
	UserInfoAddr    = "/api/user/info"
)

const (
	BUY_PRICE = iota + 1
	SELL_PRICE
	TOTAL_AMOUNT
	SINGLE_AMOUNT
)

const (
	BUY_OP = iota + 1
	SELL_OP
)

const (
	G_WEI = 10000000000
)
