package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/martian/log"
	"golang.org/x/xerrors"
	"spike-mc-ops/util"
	"strconv"
	"time"
)

func Login(privateKey string, address string, loginUrl string) (token string, err error) {
	nonce := time.Now().Unix() //1733208601
	fmt.Println("nonce: ", nonce)
	//Welcome to https://test.btc.fun, your address is 0x447d23C91E7827e6Bc8975D9999C4da2e28027de and nonce is 1733208601
	msg := "Welcome to https://test.btc.fun, your address is " + address + " and nonce is " + strconv.Itoa(int(nonce))
	signature, err := EthSign(msg, privateKey)
	if err != nil {
		log.Errorf("err: %v", err)
		return "", err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(LoginReq{
			Address:   address,
			Chain:     "EVM",
			Domain:    "https://test.btc.fun",
			Nonce:     nonce,
			Signature: signature,
		}).
		Post(loginUrl)
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}

	var loginRes LoginResp
	err = json.Unmarshal(resp.Body(), &loginRes)
	if err != nil {
		log.Errorf("err: %v, %s", err, string(resp.Body()))
		return
	}
	if loginRes.Code != "OK" {
		fmt.Println("body: ", resp.Body())
		return "", xerrors.New("code is not OK")
	}
	log.Debugf("access token: %v", loginRes.Data.Token)
	return loginRes.Data.Token, nil
}

func Sign(accessToken string, address string, amount int64, signUrl string, partyTokenContractAddress string) (res SignResp, err error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Btcfun-Login-Address", address).
		SetHeader("Btcfun-Login-Token", accessToken).
		SetBody(SignReq{
			Address: address,
			OfferOf: util.ToWei(strconv.FormatInt(2, 10), 18).String(), //只有一次募资 OfferOf为0
			Amount:  util.ToWei(strconv.FormatInt(amount, 10), 18).String(),
			Token:   partyTokenContractAddress,
		}).Post(signUrl)

	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	var signResp SignResp
	err = json.Unmarshal(resp.Body(), &signResp)
	if err != nil {
		log.Errorf("err: %v, data: %s", err, string(resp.Body()))
		return
	}
	if signResp.Code != "OK" {
		return signResp, xerrors.New("code is not OK")
	}
	log.Debugf("signResp: %v", signResp)
	return signResp, nil
}

type SignReq struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
	OfferOf string `json:"offeredOf"`
	Token   string `json:"token"`
}

type SignResp struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data SignData `json:"data"`
}

type NonceResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data NonceData `json:"data"`
}

type NonceData struct {
	Nonce string `json:"nonce"`
}

type SignData struct {
	R         string `json:"r"`
	S         string `json:"s"`
	V         int64  `json:"v"`
	Signatory string `json:"signatory"`
	Expiry    int64  `json:"expiry"`
}

type LoginReq struct {
	Address   string `json:"address"`
	Chain     string `json:"chain"`
	Domain    string `json:"domain"`
	Nonce     int64  `json:"nonce"`
	Signature string `json:"signature"`
}

type LoginResp struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Data LoginData `json:"data"`
}

type LoginData struct {
	Token string `json:"token"`
}
