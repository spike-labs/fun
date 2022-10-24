package util

import (
	"github.com/Fueav/merkletree"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type UserInfo struct {
	Address string `json:"address"`
}

func (t UserInfo) CalculateHash() ([]byte, error) {
	h := sha3.NewLegacyKeccak256()
	if _, err := h.Write(common.Hex2Bytes(t.Address)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t UserInfo) Equals(other merkletree.Content) (bool, error) {
	return t.Address == other.(UserInfo).Address, nil
}
