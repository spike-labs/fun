package config

import (
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("config")

var Cfg Config

type Config struct {
	System   System   `json:"system" toml:"system"`
	Contract Contract `json:"contract" toml:"contract"`
	Redis    Redis    `json:"redis" toml:"redis"`
	Chain    Chain    `json:"chain" toml:"chain"`
	Mysql    Mysql    `json:"mysql" toml:"mysql"`
}

type Mysql struct {
	Path         string `json:"path" toml:"path"`
	Port         string `json:"port" toml:"port"`
	Config       string `json:"config" toml:"config"`
	Dbname       string `json:"db_name" toml:"dbName"`
	Username     string `json:"username" toml:"username"`
	Password     string `json:"password" toml:"password"`
	MaxIdleConns int    `json:"maxIdleConns" toml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" toml:"maxOpenConns"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

type Chain struct {
	WsNodeAddress  string `toml:"wsNodeAddress"`
	RpcNodeAddress string `toml:"rpcNodeAddress"`
}

type System struct {
	Port      string `toml:"port"`
	MachineId string `toml:"machineId"`
}

type Redis struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
}

type Contract struct {
	SksAddress           string `toml:"SksAddress"`
	UsdcAddress          string `toml:"UsdcAddress"`
	PanCakeRouterAddress string `toml:"PanCakeRouterAddress"`
}
