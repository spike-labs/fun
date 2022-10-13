package config

import (
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("config")

var Cfg Config

type Config struct {
	System       System       `toml:"System"`
	Contract     Contract     `toml:"Contract"`
	Redis        Redis        `toml:"Redis"`
	Chain        Chain        `toml:"Chain"`
	Mysql        Mysql        `toml:"Mysql"`
	PuppetWallet PuppetWallet `toml:"PuppetWallet"`
	GameInfo     GameInfo     `toml:"Game"`
	Moralis      Moralis      `toml:"Moralis"`
}

type Moralis struct {
	XApiKey string `toml:"XApiKey"`
}

type GameInfo struct {
	DefaultPassword   string `toml:"DefaultPassword"`
	GameServerAddress string `toml:"GameServerAddr"`
	DefaultCaptcha    string `toml:"DefaultCaptcha"`
}

type Mysql struct {
	Path         string `toml:"Path"`
	Port         string `toml:"Port"`
	Config       string `toml:"Config"`
	Dbname       string `toml:"DbName"`
	Username     string `toml:"Username"`
	Password     string `toml:"Password"`
	MaxIdleConns int    `toml:"MaxIdleConns"`
	MaxOpenConns int    `toml:"MaxOpenConns"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

type Chain struct {
	WsNodeAddress  string `toml:"WsNodeAddress"`
	RpcNodeAddress string `toml:"RpcNodeAddress"`
}

type System struct {
	Port  string `toml:"Port"`
	Token string `toml:"Token"`
}

type Redis struct {
	Address  string `toml:"Address"`
	Password string `toml:"Password"`
}

type Contract struct {
	GameTokenDexPoolAddress string `toml:"GameTokenDexPoolAddress"`
	GameTokenAddress        string `toml:"GameTokenAddress"`
	UsdcAddress             string `toml:"UsdcAddress"`
	NftAddress              string `toml:"NftAddress"`
	GameVaultAddress        string `toml:"GameVaultAddress"`
	PanCakeRouterAddress    string `toml:"PanCakeRouterAddress"`
}

type PuppetWallet struct {
	PrivateKey    []string `toml:"PrivateKey"`
	Mnemonic      string   `toml:"Mnemonic"`
	AccountNumber int      `toml:"AccountNumber"`
}
