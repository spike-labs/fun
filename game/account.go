package game

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	logger "github.com/ipfs/go-log"
	"golang.org/x/xerrors"
	"spike-mc-ops/config"
	"spike-mc-ops/constant"
)

var log = logger.Logger("game")

func BindAccount(token string, address string) error {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Client-Agent", "zh-CN/CN/1/1/6/android").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Post(getBindAccountUrl(address))
	if err != nil {
		log.Error("bind account err : %+v", err)
		return err
	}
	log.Infof("resp : %s", string(resp.Body()))

	var res Resp
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return err
	}
	if res.Code != 1 {
		return xerrors.New("bind mail  error")
	}
	return nil
}

func getBindAccountUrl(account string) string {
	return fmt.Sprintf(fmt.Sprintf("%s%s?address=%s&captcha=%s", config.Cfg.GameInfo.GameServerAddress, constant.BindAccountAddr, account, config.Cfg.GameInfo.DefaultCaptcha))
}

func SendBindMail(token string) error {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Client-Agent", "zh-CN/CN/1/1/6/android").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Post(fmt.Sprintf("%s%s", config.Cfg.GameInfo.GameServerAddress, constant.SendMailAddr))
	if err != nil {
		log.Error("user info err : %+v", err)
		return err
	}
	log.Infof("resp : %s", string(resp.Body()))

	var res Resp
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return err
	}
	if res.Code != 1 {
		return xerrors.New("send bind mail  error")
	}
	return nil
}

func QueryUserInfo(token string) error {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Client-Agent", "zh-CN/CN/1/1/6/android").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		Post(fmt.Sprintf("%s%s", config.Cfg.GameInfo.GameServerAddress, constant.UserInfoAddr))
	if err != nil {
		log.Error("user info err : %+v", err)
		return err
	}
	var res Resp
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return err
	}
	if res.Code != 1 {
		return xerrors.New("query user info error")
	}
	return nil
}

func Login(username string) (token string, err error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Client-Agent", "zh-CN/CN/1/1/6/android").
		Post(getLoginUrl(username))
	if err != nil {
		log.Error("user login err : %+v", err)
		return "", err
	}
	var res loginResp
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return "", err
	}
	log.Infof("token : %s", res.Data.Token)
	return res.Data.Token, nil
}

type Resp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type loginResp struct {
	Code int64     `json:"code"`
	Msg  string    `json:"msg"`
	Data loginData `json:"data"`
}

type loginData struct {
	Token string `json:"tokenKey"`
}

func getLoginUrl(username string) string {
	return fmt.Sprintf("%s%s?account=%s&password=%s", config.Cfg.GameInfo.GameServerAddress, constant.LoginAddr, username, config.Cfg.GameInfo.DefaultPassword)
}
