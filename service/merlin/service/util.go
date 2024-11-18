package service

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/martian/log"
)

func EthSign(message string, privateKeyHex string) (sig string, err error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		log.Errorf("Failed to parse private key: %v", err)
		return "", err
	}
	hashedMessage := TextHash([]byte(message))

	signature, err := crypto.Sign(hashedMessage, privateKey)
	if err != nil {
		log.Errorf("Failed to sign message: %v", err)
		return "", err
	}
	return fmt.Sprintf("0x%x", signature), nil
}

func TextHash(message []byte) []byte {
	prefix := []byte("\x19Ethereum Signed Message:\n")
	length := len(message)

	messageLength := []byte(fmt.Sprintf("%d", length))

	fullMessage := append(prefix, append(messageLength, message...)...)

	hash := crypto.Keccak256(fullMessage)
	return hash
}

func HexToByte32(hexStr string) ([32]byte, error) {
	if len(hexStr) != 64 {
		return [32]byte{}, fmt.Errorf("invalid hex string length: expected 64 characters, got %d", len(hexStr))
	}

	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to decode hex string: %v", err)
	}

	var byteArray [32]byte
	copy(byteArray[:], bytes)
	return byteArray, nil
}
