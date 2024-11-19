package service

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/google/martian/log"
	"spike-mc-ops/util"
	"strconv"
)

func Login(privateKey string, address string, nonce string, loginUrl string) (token string, err error) {
	msg := "Sign in BTC.FUN\n\nWallet address:\n" + address + "\n\nExpiration time:\n2025-01-27T17:19:13+08:00\nNonce:\n" + nonce + "\n"
	signature, err := EthSign(msg, privateKey)
	if err != nil {
		log.Errorf("err: %v", err)
		return "", err
	}
	log.Debugf("signature: %s", signature)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(LoginReq{
			Type:      "evm",
			Message:   msg,
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
	log.Debugf("access token: %v", loginRes.Data.AccessToken)
	return loginRes.Data.AccessToken, nil
}

func QueryNonce(address string, nonceUrl string) (res NonceResp, err error) {
	log.Debugf("QueryNonce, address: %s", address)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(nonceUrl + address)

	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	var nonceRes NonceResp
	err = json.Unmarshal(resp.Body(), &nonceRes)
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	return nonceRes, nil
}

func Sign(accessToken string, address string, amount int64, signUrl string, partyTokenContractAddress string) (res SignResp, err error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+accessToken).
		SetBody(SignReq{
			Address: address,
			OfferOf: util.ToWei(strconv.FormatInt(0, 10), 18).String(), //只有一次募资 OfferOf为0
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
	Code int      `json:"code"`
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
	TimeStamp int64  `json:"timeStamp"`
}

type LoginReq struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Type      string `json:"type"`
}

type LoginResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data LoginData `json:"data"`
}

type LoginData struct {
	AccessToken string `json:"access_token"`
}
