package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"regexp"
	"spike-mc-ops/model"
	"strconv"
)

func ValidateAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func NewMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Error("NewEntropy", err)
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Error("NewMnemonic", err)
		return mnemonic, err
	}

	log.Infof(mnemonic)
	return mnemonic, nil
}

func DerivePrivateKeyWithNumber(mnemonic string, number int) (string, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Error("NewFromMnemonic", err)
		return "", err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.Itoa(number))
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Error("Derive", err)
		return "", err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		log.Error("Private Key", err)
		return "", err
	}

	return hexutil.Encode(crypto.FromECDSA(privateKey)), nil
}

func DerivePrivateKey(mnemonic string) (string, error) {
	return DerivePrivateKeyWithNumber(mnemonic, 0)
}

func GenerateAddress(private string) (string, error) {
	privateKey, err := crypto.HexToECDSA(private[2:])
	if err != nil {
		log.Error("Load ECDSA from private key", err)
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("Load ECDSA Public Key", err)
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).String()

	return address, err
}

func GenerateAccount(private string) (model.Account, error) {
	privateKey, err := crypto.HexToECDSA(private[2:])
	if err != nil {
		log.Error("Load ECDSA from private key", err)
		return model.Account{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("Load ECDSA Public Key", err)
		return model.Account{}, err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).String()

	return model.Account{
		PrivateKey: private,
		Address:    address,
	}, nil
}
