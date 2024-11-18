// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IiZiSwapPoolLimitOrderStruct is an auto generated low-level Go binding around an user-defined struct.
type IiZiSwapPoolLimitOrderStruct struct {
	SellingX *big.Int
	EarnY    *big.Int
	AccEarnY *big.Int
	SellingY *big.Int
	EarnX    *big.Int
	AccEarnX *big.Int
}

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"addAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"acquireAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimSold\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimEarn\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"AddLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"collectDec\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"collectEarn\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"CollectLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"CollectLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"decreaseAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimSold\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"claimEarn\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"}],\"name\":\"DecLimitOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidY\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sellXEarnY\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amountX\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addLimOrderWithX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"orderX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"acquireY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amountY\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addLimOrderWithY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"orderY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"acquireX\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"assignX\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"fromLegacy\",\"type\":\"bool\"}],\"name\":\"assignLimOrderEarnX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualAssignX\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"assignY\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"fromLegacy\",\"type\":\"bool\"}],\"name\":\"assignLimOrderEarnY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualAssignY\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidDelta\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amountXLim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYLim\",\"type\":\"uint256\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualAmountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualAmountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFeeCharged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"collectDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"collectEarn\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isEarnY\",\"type\":\"bool\"}],\"name\":\"collectLimOrder\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualCollectDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"actualCollectEarn\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"deltaX\",\"type\":\"uint128\"}],\"name\":\"decLimOrderWithX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualDeltaX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"point\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"deltaY\",\"type\":\"uint128\"}],\"name\":\"decLimOrderWithY\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"actualDeltaY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newNextQueueLen\",\"type\":\"uint16\"}],\"name\":\"expandObservationQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeChargePercent\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeScaleX_128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeScaleY_128\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leftMostPt\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"limitOrderData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"sellingX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarnY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarnX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"legacyAccEarnX\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"}],\"name\":\"limitOrderSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"sellingX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnY\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnY\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingY\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnX\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accEarnX\",\"type\":\"uint256\"}],\"internalType\":\"structIiZiSwapPool.LimitOrderStruct[]\",\"name\":\"limitOrders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleX_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleY_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedY\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"leftPoint\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPoint\",\"type\":\"int24\"}],\"name\":\"liquiditySnapshot\",\"outputs\":[{\"internalType\":\"int128[]\",\"name\":\"deltaLiquidities\",\"type\":\"int128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidPt\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidDelta\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newFeeChargePercent\",\"type\":\"uint24\"}],\"name\":\"modifyFeeChargePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"observations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"accPoint\",\"type\":\"int56\"},{\"internalType\":\"bool\",\"name\":\"init\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"observe\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"accPoints\",\"type\":\"int56[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"orderOrEndpoint\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"name\":\"pointBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pointDelta\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"points\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidSum\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidDelta\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"accFeeXOut_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accFeeYOut_128\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEndpt\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rightMostPt\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sqrtRate_96\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPrice_96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationCurrentIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationNextQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityX\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desireY\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2X\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desireX\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeXCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeYCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"userEarnX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastAccEarn\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingRemain\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnAssign\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"userEarnY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastAccEarn\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"sellingRemain\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sellingDec\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"legacyEarn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"earnAssign\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Swap *SwapCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Swap *SwapSession) Factory() (common.Address, error) {
	return _Swap.Contract.Factory(&_Swap.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Swap *SwapCallerSession) Factory() (common.Address, error) {
	return _Swap.Contract.Factory(&_Swap.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Swap *SwapCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Swap *SwapSession) Fee() (*big.Int, error) {
	return _Swap.Contract.Fee(&_Swap.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Swap *SwapCallerSession) Fee() (*big.Int, error) {
	return _Swap.Contract.Fee(&_Swap.CallOpts)
}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_Swap *SwapCaller) FeeChargePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "feeChargePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_Swap *SwapSession) FeeChargePercent() (*big.Int, error) {
	return _Swap.Contract.FeeChargePercent(&_Swap.CallOpts)
}

// FeeChargePercent is a free data retrieval call binding the contract method 0x81794fba.
//
// Solidity: function feeChargePercent() view returns(uint24)
func (_Swap *SwapCallerSession) FeeChargePercent() (*big.Int, error) {
	return _Swap.Contract.FeeChargePercent(&_Swap.CallOpts)
}

// FeeScaleX128 is a free data retrieval call binding the contract method 0x1aae2e55.
//
// Solidity: function feeScaleX_128() view returns(uint256)
func (_Swap *SwapCaller) FeeScaleX128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "feeScaleX_128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeScaleX128 is a free data retrieval call binding the contract method 0x1aae2e55.
//
// Solidity: function feeScaleX_128() view returns(uint256)
func (_Swap *SwapSession) FeeScaleX128() (*big.Int, error) {
	return _Swap.Contract.FeeScaleX128(&_Swap.CallOpts)
}

// FeeScaleX128 is a free data retrieval call binding the contract method 0x1aae2e55.
//
// Solidity: function feeScaleX_128() view returns(uint256)
func (_Swap *SwapCallerSession) FeeScaleX128() (*big.Int, error) {
	return _Swap.Contract.FeeScaleX128(&_Swap.CallOpts)
}

// FeeScaleY128 is a free data retrieval call binding the contract method 0x588e59ae.
//
// Solidity: function feeScaleY_128() view returns(uint256)
func (_Swap *SwapCaller) FeeScaleY128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "feeScaleY_128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeScaleY128 is a free data retrieval call binding the contract method 0x588e59ae.
//
// Solidity: function feeScaleY_128() view returns(uint256)
func (_Swap *SwapSession) FeeScaleY128() (*big.Int, error) {
	return _Swap.Contract.FeeScaleY128(&_Swap.CallOpts)
}

// FeeScaleY128 is a free data retrieval call binding the contract method 0x588e59ae.
//
// Solidity: function feeScaleY_128() view returns(uint256)
func (_Swap *SwapCallerSession) FeeScaleY128() (*big.Int, error) {
	return _Swap.Contract.FeeScaleY128(&_Swap.CallOpts)
}

// LeftMostPt is a free data retrieval call binding the contract method 0x537c2d8e.
//
// Solidity: function leftMostPt() view returns(int24)
func (_Swap *SwapCaller) LeftMostPt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "leftMostPt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeftMostPt is a free data retrieval call binding the contract method 0x537c2d8e.
//
// Solidity: function leftMostPt() view returns(int24)
func (_Swap *SwapSession) LeftMostPt() (*big.Int, error) {
	return _Swap.Contract.LeftMostPt(&_Swap.CallOpts)
}

// LeftMostPt is a free data retrieval call binding the contract method 0x537c2d8e.
//
// Solidity: function leftMostPt() view returns(int24)
func (_Swap *SwapCallerSession) LeftMostPt() (*big.Int, error) {
	return _Swap.Contract.LeftMostPt(&_Swap.CallOpts)
}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 ) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint256 legacyAccEarnY, uint128 legacyEarnY, uint128 sellingY, uint128 earnX, uint128 legacyEarnX, uint256 accEarnX, uint256 legacyAccEarnX)
func (_Swap *SwapCaller) LimitOrderData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SellingX       *big.Int
	EarnY          *big.Int
	AccEarnY       *big.Int
	LegacyAccEarnY *big.Int
	LegacyEarnY    *big.Int
	SellingY       *big.Int
	EarnX          *big.Int
	LegacyEarnX    *big.Int
	AccEarnX       *big.Int
	LegacyAccEarnX *big.Int
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "limitOrderData", arg0)

	outstruct := new(struct {
		SellingX       *big.Int
		EarnY          *big.Int
		AccEarnY       *big.Int
		LegacyAccEarnY *big.Int
		LegacyEarnY    *big.Int
		SellingY       *big.Int
		EarnX          *big.Int
		LegacyEarnX    *big.Int
		AccEarnX       *big.Int
		LegacyAccEarnX *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SellingX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EarnY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccEarnY = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LegacyAccEarnY = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarnY = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SellingY = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EarnX = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarnX = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.AccEarnX = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LegacyAccEarnX = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 ) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint256 legacyAccEarnY, uint128 legacyEarnY, uint128 sellingY, uint128 earnX, uint128 legacyEarnX, uint256 accEarnX, uint256 legacyAccEarnX)
func (_Swap *SwapSession) LimitOrderData(arg0 *big.Int) (struct {
	SellingX       *big.Int
	EarnY          *big.Int
	AccEarnY       *big.Int
	LegacyAccEarnY *big.Int
	LegacyEarnY    *big.Int
	SellingY       *big.Int
	EarnX          *big.Int
	LegacyEarnX    *big.Int
	AccEarnX       *big.Int
	LegacyAccEarnX *big.Int
}, error) {
	return _Swap.Contract.LimitOrderData(&_Swap.CallOpts, arg0)
}

// LimitOrderData is a free data retrieval call binding the contract method 0x8790aca3.
//
// Solidity: function limitOrderData(int24 ) view returns(uint128 sellingX, uint128 earnY, uint256 accEarnY, uint256 legacyAccEarnY, uint128 legacyEarnY, uint128 sellingY, uint128 earnX, uint128 legacyEarnX, uint256 accEarnX, uint256 legacyAccEarnX)
func (_Swap *SwapCallerSession) LimitOrderData(arg0 *big.Int) (struct {
	SellingX       *big.Int
	EarnY          *big.Int
	AccEarnY       *big.Int
	LegacyAccEarnY *big.Int
	LegacyEarnY    *big.Int
	SellingY       *big.Int
	EarnX          *big.Int
	LegacyEarnX    *big.Int
	AccEarnX       *big.Int
	LegacyAccEarnX *big.Int
}, error) {
	return _Swap.Contract.LimitOrderData(&_Swap.CallOpts, arg0)
}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_Swap *SwapCaller) LimitOrderSnapshot(opts *bind.CallOpts, leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "limitOrderSnapshot", leftPoint, rightPoint)

	if err != nil {
		return *new([]IiZiSwapPoolLimitOrderStruct), err
	}

	out0 := *abi.ConvertType(out[0], new([]IiZiSwapPoolLimitOrderStruct)).(*[]IiZiSwapPoolLimitOrderStruct)

	return out0, err

}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_Swap *SwapSession) LimitOrderSnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	return _Swap.Contract.LimitOrderSnapshot(&_Swap.CallOpts, leftPoint, rightPoint)
}

// LimitOrderSnapshot is a free data retrieval call binding the contract method 0x6f73f006.
//
// Solidity: function limitOrderSnapshot(int24 leftPoint, int24 rightPoint) view returns((uint128,uint128,uint256,uint128,uint128,uint256)[] limitOrders)
func (_Swap *SwapCallerSession) LimitOrderSnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]IiZiSwapPoolLimitOrderStruct, error) {
	return _Swap.Contract.LimitOrderSnapshot(&_Swap.CallOpts, leftPoint, rightPoint)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 ) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_Swap *SwapCaller) Liquidity(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "liquidity", arg0)

	outstruct := new(struct {
		Liquidity        *big.Int
		LastFeeScaleX128 *big.Int
		LastFeeScaleY128 *big.Int
		TokenOwedX       *big.Int
		TokenOwedY       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleY128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedX = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedY = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 ) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_Swap *SwapSession) Liquidity(arg0 [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _Swap.Contract.Liquidity(&_Swap.CallOpts, arg0)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 ) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_Swap *SwapCallerSession) Liquidity(arg0 [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _Swap.Contract.Liquidity(&_Swap.CallOpts, arg0)
}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_Swap *SwapCaller) LiquiditySnapshot(opts *bind.CallOpts, leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "liquiditySnapshot", leftPoint, rightPoint)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_Swap *SwapSession) LiquiditySnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	return _Swap.Contract.LiquiditySnapshot(&_Swap.CallOpts, leftPoint, rightPoint)
}

// LiquiditySnapshot is a free data retrieval call binding the contract method 0xb14184e6.
//
// Solidity: function liquiditySnapshot(int24 leftPoint, int24 rightPoint) view returns(int128[] deltaLiquidities)
func (_Swap *SwapCallerSession) LiquiditySnapshot(leftPoint *big.Int, rightPoint *big.Int) ([]*big.Int, error) {
	return _Swap.Contract.LiquiditySnapshot(&_Swap.CallOpts, leftPoint, rightPoint)
}

// MaxLiquidPt is a free data retrieval call binding the contract method 0x6d01843b.
//
// Solidity: function maxLiquidPt() view returns(uint128)
func (_Swap *SwapCaller) MaxLiquidPt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "maxLiquidPt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidPt is a free data retrieval call binding the contract method 0x6d01843b.
//
// Solidity: function maxLiquidPt() view returns(uint128)
func (_Swap *SwapSession) MaxLiquidPt() (*big.Int, error) {
	return _Swap.Contract.MaxLiquidPt(&_Swap.CallOpts)
}

// MaxLiquidPt is a free data retrieval call binding the contract method 0x6d01843b.
//
// Solidity: function maxLiquidPt() view returns(uint128)
func (_Swap *SwapCallerSession) MaxLiquidPt() (*big.Int, error) {
	return _Swap.Contract.MaxLiquidPt(&_Swap.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_Swap *SwapCaller) Observations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "observations", arg0)

	outstruct := new(struct {
		Timestamp uint32
		AccPoint  *big.Int
		Init      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.AccPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Init = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_Swap *SwapSession) Observations(arg0 *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	return _Swap.Contract.Observations(&_Swap.CallOpts, arg0)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 ) view returns(uint32 timestamp, int56 accPoint, bool init)
func (_Swap *SwapCallerSession) Observations(arg0 *big.Int) (struct {
	Timestamp uint32
	AccPoint  *big.Int
	Init      bool
}, error) {
	return _Swap.Contract.Observations(&_Swap.CallOpts, arg0)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_Swap *SwapCaller) Observe(opts *bind.CallOpts, secondsAgos []uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "observe", secondsAgos)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_Swap *SwapSession) Observe(secondsAgos []uint32) ([]*big.Int, error) {
	return _Swap.Contract.Observe(&_Swap.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] accPoints)
func (_Swap *SwapCallerSession) Observe(secondsAgos []uint32) ([]*big.Int, error) {
	return _Swap.Contract.Observe(&_Swap.CallOpts, secondsAgos)
}

// OrderOrEndpoint is a free data retrieval call binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 ) view returns(int24)
func (_Swap *SwapCaller) OrderOrEndpoint(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "orderOrEndpoint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderOrEndpoint is a free data retrieval call binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 ) view returns(int24)
func (_Swap *SwapSession) OrderOrEndpoint(arg0 *big.Int) (*big.Int, error) {
	return _Swap.Contract.OrderOrEndpoint(&_Swap.CallOpts, arg0)
}

// OrderOrEndpoint is a free data retrieval call binding the contract method 0xedcba3b2.
//
// Solidity: function orderOrEndpoint(int24 ) view returns(int24)
func (_Swap *SwapCallerSession) OrderOrEndpoint(arg0 *big.Int) (*big.Int, error) {
	return _Swap.Contract.OrderOrEndpoint(&_Swap.CallOpts, arg0)
}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 ) view returns(uint256)
func (_Swap *SwapCaller) PointBitmap(opts *bind.CallOpts, arg0 int16) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "pointBitmap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 ) view returns(uint256)
func (_Swap *SwapSession) PointBitmap(arg0 int16) (*big.Int, error) {
	return _Swap.Contract.PointBitmap(&_Swap.CallOpts, arg0)
}

// PointBitmap is a free data retrieval call binding the contract method 0x98a0f72e.
//
// Solidity: function pointBitmap(int16 ) view returns(uint256)
func (_Swap *SwapCallerSession) PointBitmap(arg0 int16) (*big.Int, error) {
	return _Swap.Contract.PointBitmap(&_Swap.CallOpts, arg0)
}

// PointDelta is a free data retrieval call binding the contract method 0x58c51ce6.
//
// Solidity: function pointDelta() view returns(int24)
func (_Swap *SwapCaller) PointDelta(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "pointDelta")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointDelta is a free data retrieval call binding the contract method 0x58c51ce6.
//
// Solidity: function pointDelta() view returns(int24)
func (_Swap *SwapSession) PointDelta() (*big.Int, error) {
	return _Swap.Contract.PointDelta(&_Swap.CallOpts)
}

// PointDelta is a free data retrieval call binding the contract method 0x58c51ce6.
//
// Solidity: function pointDelta() view returns(int24)
func (_Swap *SwapCallerSession) PointDelta() (*big.Int, error) {
	return _Swap.Contract.PointDelta(&_Swap.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 ) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_Swap *SwapCaller) Points(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "points", arg0)

	outstruct := new(struct {
		LiquidSum     *big.Int
		LiquidDelta   *big.Int
		AccFeeXOut128 *big.Int
		AccFeeYOut128 *big.Int
		IsEndpt       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidSum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccFeeXOut128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccFeeYOut128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsEndpt = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 ) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_Swap *SwapSession) Points(arg0 *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	return _Swap.Contract.Points(&_Swap.CallOpts, arg0)
}

// Points is a free data retrieval call binding the contract method 0x75c0e0d5.
//
// Solidity: function points(int24 ) view returns(uint128 liquidSum, int128 liquidDelta, uint256 accFeeXOut_128, uint256 accFeeYOut_128, bool isEndpt)
func (_Swap *SwapCallerSession) Points(arg0 *big.Int) (struct {
	LiquidSum     *big.Int
	LiquidDelta   *big.Int
	AccFeeXOut128 *big.Int
	AccFeeYOut128 *big.Int
	IsEndpt       bool
}, error) {
	return _Swap.Contract.Points(&_Swap.CallOpts, arg0)
}

// RightMostPt is a free data retrieval call binding the contract method 0xd3b16864.
//
// Solidity: function rightMostPt() view returns(int24)
func (_Swap *SwapCaller) RightMostPt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "rightMostPt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RightMostPt is a free data retrieval call binding the contract method 0xd3b16864.
//
// Solidity: function rightMostPt() view returns(int24)
func (_Swap *SwapSession) RightMostPt() (*big.Int, error) {
	return _Swap.Contract.RightMostPt(&_Swap.CallOpts)
}

// RightMostPt is a free data retrieval call binding the contract method 0xd3b16864.
//
// Solidity: function rightMostPt() view returns(int24)
func (_Swap *SwapCallerSession) RightMostPt() (*big.Int, error) {
	return _Swap.Contract.RightMostPt(&_Swap.CallOpts)
}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_Swap *SwapCaller) SqrtRate96(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "sqrtRate_96")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_Swap *SwapSession) SqrtRate96() (*big.Int, error) {
	return _Swap.Contract.SqrtRate96(&_Swap.CallOpts)
}

// SqrtRate96 is a free data retrieval call binding the contract method 0x09beabc1.
//
// Solidity: function sqrtRate_96() view returns(uint160)
func (_Swap *SwapCallerSession) SqrtRate96() (*big.Int, error) {
	return _Swap.Contract.SqrtRate96(&_Swap.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_Swap *SwapCaller) State(opts *bind.CallOpts) (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		SqrtPrice96             *big.Int
		CurrentPoint            *big.Int
		ObservationCurrentIndex uint16
		ObservationQueueLen     uint16
		ObservationNextQueueLen uint16
		Locked                  bool
		Liquidity               *big.Int
		LiquidityX              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPrice96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationCurrentIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationQueueLen = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationNextQueueLen = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Locked = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Liquidity = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LiquidityX = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_Swap *SwapSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _Swap.Contract.State(&_Swap.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_Swap *SwapCallerSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _Swap.Contract.State(&_Swap.CallOpts)
}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Swap *SwapCaller) TokenX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "tokenX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Swap *SwapSession) TokenX() (common.Address, error) {
	return _Swap.Contract.TokenX(&_Swap.CallOpts)
}

// TokenX is a free data retrieval call binding the contract method 0x16dc165b.
//
// Solidity: function tokenX() view returns(address)
func (_Swap *SwapCallerSession) TokenX() (common.Address, error) {
	return _Swap.Contract.TokenX(&_Swap.CallOpts)
}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Swap *SwapCaller) TokenY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "tokenY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Swap *SwapSession) TokenY() (common.Address, error) {
	return _Swap.Contract.TokenY(&_Swap.CallOpts)
}

// TokenY is a free data retrieval call binding the contract method 0xb7d19fc4.
//
// Solidity: function tokenY() view returns(address)
func (_Swap *SwapCallerSession) TokenY() (common.Address, error) {
	return _Swap.Contract.TokenY(&_Swap.CallOpts)
}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_Swap *SwapCaller) TotalFeeXCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "totalFeeXCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_Swap *SwapSession) TotalFeeXCharged() (*big.Int, error) {
	return _Swap.Contract.TotalFeeXCharged(&_Swap.CallOpts)
}

// TotalFeeXCharged is a free data retrieval call binding the contract method 0xe556289f.
//
// Solidity: function totalFeeXCharged() view returns(uint256)
func (_Swap *SwapCallerSession) TotalFeeXCharged() (*big.Int, error) {
	return _Swap.Contract.TotalFeeXCharged(&_Swap.CallOpts)
}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_Swap *SwapCaller) TotalFeeYCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "totalFeeYCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_Swap *SwapSession) TotalFeeYCharged() (*big.Int, error) {
	return _Swap.Contract.TotalFeeYCharged(&_Swap.CallOpts)
}

// TotalFeeYCharged is a free data retrieval call binding the contract method 0x33005cd5.
//
// Solidity: function totalFeeYCharged() view returns(uint256)
func (_Swap *SwapCallerSession) TotalFeeYCharged() (*big.Int, error) {
	return _Swap.Contract.TotalFeeYCharged(&_Swap.CallOpts)
}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_Swap *SwapCaller) UserEarnX(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "userEarnX", arg0)

	outstruct := new(struct {
		LastAccEarn   *big.Int
		SellingRemain *big.Int
		SellingDec    *big.Int
		Earn          *big.Int
		LegacyEarn    *big.Int
		EarnAssign    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastAccEarn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SellingRemain = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellingDec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Earn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarn = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EarnAssign = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_Swap *SwapSession) UserEarnX(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Swap.Contract.UserEarnX(&_Swap.CallOpts, arg0)
}

// UserEarnX is a free data retrieval call binding the contract method 0x62ccaafd.
//
// Solidity: function userEarnX(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_Swap *SwapCallerSession) UserEarnX(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Swap.Contract.UserEarnX(&_Swap.CallOpts, arg0)
}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_Swap *SwapCaller) UserEarnY(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "userEarnY", arg0)

	outstruct := new(struct {
		LastAccEarn   *big.Int
		SellingRemain *big.Int
		SellingDec    *big.Int
		Earn          *big.Int
		LegacyEarn    *big.Int
		EarnAssign    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastAccEarn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SellingRemain = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellingDec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Earn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LegacyEarn = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EarnAssign = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_Swap *SwapSession) UserEarnY(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Swap.Contract.UserEarnY(&_Swap.CallOpts, arg0)
}

// UserEarnY is a free data retrieval call binding the contract method 0x1621835f.
//
// Solidity: function userEarnY(bytes32 ) view returns(uint256 lastAccEarn, uint128 sellingRemain, uint128 sellingDec, uint128 earn, uint128 legacyEarn, uint128 earnAssign)
func (_Swap *SwapCallerSession) UserEarnY(arg0 [32]byte) (struct {
	LastAccEarn   *big.Int
	SellingRemain *big.Int
	SellingDec    *big.Int
	Earn          *big.Int
	LegacyEarn    *big.Int
	EarnAssign    *big.Int
}, error) {
	return _Swap.Contract.UserEarnY(&_Swap.CallOpts, arg0)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_Swap *SwapTransactor) AddLimOrderWithX(opts *bind.TransactOpts, recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "addLimOrderWithX", recipient, point, amountX, data)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_Swap *SwapSession) AddLimOrderWithX(recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.AddLimOrderWithX(&_Swap.TransactOpts, recipient, point, amountX, data)
}

// AddLimOrderWithX is a paid mutator transaction binding the contract method 0xff12504e.
//
// Solidity: function addLimOrderWithX(address recipient, int24 point, uint128 amountX, bytes data) returns(uint128 orderX, uint128 acquireY)
func (_Swap *SwapTransactorSession) AddLimOrderWithX(recipient common.Address, point *big.Int, amountX *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.AddLimOrderWithX(&_Swap.TransactOpts, recipient, point, amountX, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_Swap *SwapTransactor) AddLimOrderWithY(opts *bind.TransactOpts, recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "addLimOrderWithY", recipient, point, amountY, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_Swap *SwapSession) AddLimOrderWithY(recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.AddLimOrderWithY(&_Swap.TransactOpts, recipient, point, amountY, data)
}

// AddLimOrderWithY is a paid mutator transaction binding the contract method 0x0e1552f0.
//
// Solidity: function addLimOrderWithY(address recipient, int24 point, uint128 amountY, bytes data) returns(uint128 orderY, uint128 acquireX)
func (_Swap *SwapTransactorSession) AddLimOrderWithY(recipient common.Address, point *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.AddLimOrderWithY(&_Swap.TransactOpts, recipient, point, amountY, data)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xf0163ef4.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX, bool fromLegacy) returns(uint128 actualAssignX)
func (_Swap *SwapTransactor) AssignLimOrderEarnX(opts *bind.TransactOpts, point *big.Int, assignX *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "assignLimOrderEarnX", point, assignX, fromLegacy)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xf0163ef4.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX, bool fromLegacy) returns(uint128 actualAssignX)
func (_Swap *SwapSession) AssignLimOrderEarnX(point *big.Int, assignX *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _Swap.Contract.AssignLimOrderEarnX(&_Swap.TransactOpts, point, assignX, fromLegacy)
}

// AssignLimOrderEarnX is a paid mutator transaction binding the contract method 0xf0163ef4.
//
// Solidity: function assignLimOrderEarnX(int24 point, uint128 assignX, bool fromLegacy) returns(uint128 actualAssignX)
func (_Swap *SwapTransactorSession) AssignLimOrderEarnX(point *big.Int, assignX *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _Swap.Contract.AssignLimOrderEarnX(&_Swap.TransactOpts, point, assignX, fromLegacy)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0x544e7057.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY, bool fromLegacy) returns(uint128 actualAssignY)
func (_Swap *SwapTransactor) AssignLimOrderEarnY(opts *bind.TransactOpts, point *big.Int, assignY *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "assignLimOrderEarnY", point, assignY, fromLegacy)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0x544e7057.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY, bool fromLegacy) returns(uint128 actualAssignY)
func (_Swap *SwapSession) AssignLimOrderEarnY(point *big.Int, assignY *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _Swap.Contract.AssignLimOrderEarnY(&_Swap.TransactOpts, point, assignY, fromLegacy)
}

// AssignLimOrderEarnY is a paid mutator transaction binding the contract method 0x544e7057.
//
// Solidity: function assignLimOrderEarnY(int24 point, uint128 assignY, bool fromLegacy) returns(uint128 actualAssignY)
func (_Swap *SwapTransactorSession) AssignLimOrderEarnY(point *big.Int, assignY *big.Int, fromLegacy bool) (*types.Transaction, error) {
	return _Swap.Contract.AssignLimOrderEarnY(&_Swap.TransactOpts, point, assignY, fromLegacy)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactor) Burn(opts *bind.TransactOpts, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "burn", leftPt, rightPt, liquidDelta)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapSession) Burn(leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Burn(&_Swap.TransactOpts, leftPt, rightPt, liquidDelta)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 leftPt, int24 rightPt, uint128 liquidDelta) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactorSession) Burn(leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Burn(&_Swap.TransactOpts, leftPt, rightPt, liquidDelta)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_Swap *SwapTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "collect", recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_Swap *SwapSession) Collect(recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Collect(&_Swap.TransactOpts, recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0x872d1f15.
//
// Solidity: function collect(address recipient, int24 leftPt, int24 rightPt, uint256 amountXLim, uint256 amountYLim) returns(uint256 actualAmountX, uint256 actualAmountY)
func (_Swap *SwapTransactorSession) Collect(recipient common.Address, leftPt *big.Int, rightPt *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Collect(&_Swap.TransactOpts, recipient, leftPt, rightPt, amountXLim, amountYLim)
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_Swap *SwapTransactor) CollectFeeCharged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "collectFeeCharged")
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_Swap *SwapSession) CollectFeeCharged() (*types.Transaction, error) {
	return _Swap.Contract.CollectFeeCharged(&_Swap.TransactOpts)
}

// CollectFeeCharged is a paid mutator transaction binding the contract method 0xb74d60a9.
//
// Solidity: function collectFeeCharged() returns()
func (_Swap *SwapTransactorSession) CollectFeeCharged() (*types.Transaction, error) {
	return _Swap.Contract.CollectFeeCharged(&_Swap.TransactOpts)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_Swap *SwapTransactor) CollectLimOrder(opts *bind.TransactOpts, recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "collectLimOrder", recipient, point, collectDec, collectEarn, isEarnY)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_Swap *SwapSession) CollectLimOrder(recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _Swap.Contract.CollectLimOrder(&_Swap.TransactOpts, recipient, point, collectDec, collectEarn, isEarnY)
}

// CollectLimOrder is a paid mutator transaction binding the contract method 0x6ad1718f.
//
// Solidity: function collectLimOrder(address recipient, int24 point, uint128 collectDec, uint128 collectEarn, bool isEarnY) returns(uint128 actualCollectDec, uint128 actualCollectEarn)
func (_Swap *SwapTransactorSession) CollectLimOrder(recipient common.Address, point *big.Int, collectDec *big.Int, collectEarn *big.Int, isEarnY bool) (*types.Transaction, error) {
	return _Swap.Contract.CollectLimOrder(&_Swap.TransactOpts, recipient, point, collectDec, collectEarn, isEarnY)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX, uint256 legacyAccEarn)
func (_Swap *SwapTransactor) DecLimOrderWithX(opts *bind.TransactOpts, point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "decLimOrderWithX", point, deltaX)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX, uint256 legacyAccEarn)
func (_Swap *SwapSession) DecLimOrderWithX(point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.DecLimOrderWithX(&_Swap.TransactOpts, point, deltaX)
}

// DecLimOrderWithX is a paid mutator transaction binding the contract method 0x4cd70e91.
//
// Solidity: function decLimOrderWithX(int24 point, uint128 deltaX) returns(uint128 actualDeltaX, uint256 legacyAccEarn)
func (_Swap *SwapTransactorSession) DecLimOrderWithX(point *big.Int, deltaX *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.DecLimOrderWithX(&_Swap.TransactOpts, point, deltaX)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY, uint256 legacyAccEarn)
func (_Swap *SwapTransactor) DecLimOrderWithY(opts *bind.TransactOpts, point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "decLimOrderWithY", point, deltaY)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY, uint256 legacyAccEarn)
func (_Swap *SwapSession) DecLimOrderWithY(point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.DecLimOrderWithY(&_Swap.TransactOpts, point, deltaY)
}

// DecLimOrderWithY is a paid mutator transaction binding the contract method 0x62c944ca.
//
// Solidity: function decLimOrderWithY(int24 point, uint128 deltaY) returns(uint128 actualDeltaY, uint256 legacyAccEarn)
func (_Swap *SwapTransactorSession) DecLimOrderWithY(point *big.Int, deltaY *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.DecLimOrderWithY(&_Swap.TransactOpts, point, deltaY)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_Swap *SwapTransactor) ExpandObservationQueue(opts *bind.TransactOpts, newNextQueueLen uint16) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "expandObservationQueue", newNextQueueLen)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_Swap *SwapSession) ExpandObservationQueue(newNextQueueLen uint16) (*types.Transaction, error) {
	return _Swap.Contract.ExpandObservationQueue(&_Swap.TransactOpts, newNextQueueLen)
}

// ExpandObservationQueue is a paid mutator transaction binding the contract method 0x17fdacb9.
//
// Solidity: function expandObservationQueue(uint16 newNextQueueLen) returns()
func (_Swap *SwapTransactorSession) ExpandObservationQueue(newNextQueueLen uint16) (*types.Transaction, error) {
	return _Swap.Contract.ExpandObservationQueue(&_Swap.TransactOpts, newNextQueueLen)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_Swap *SwapTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "flash", recipient, amountX, amountY, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_Swap *SwapSession) Flash(recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.Flash(&_Swap.TransactOpts, recipient, amountX, amountY, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amountX, uint256 amountY, bytes data) returns()
func (_Swap *SwapTransactorSession) Flash(recipient common.Address, amountX *big.Int, amountY *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.Flash(&_Swap.TransactOpts, recipient, amountX, amountY, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "mint", recipient, leftPt, rightPt, liquidDelta, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapSession) Mint(recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.Mint(&_Swap.TransactOpts, recipient, leftPt, rightPt, liquidDelta, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 leftPt, int24 rightPt, uint128 liquidDelta, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactorSession) Mint(recipient common.Address, leftPt *big.Int, rightPt *big.Int, liquidDelta *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.Mint(&_Swap.TransactOpts, recipient, leftPt, rightPt, liquidDelta, data)
}

// ModifyFeeChargePercent is a paid mutator transaction binding the contract method 0xd38a85ad.
//
// Solidity: function modifyFeeChargePercent(uint24 newFeeChargePercent) returns()
func (_Swap *SwapTransactor) ModifyFeeChargePercent(opts *bind.TransactOpts, newFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "modifyFeeChargePercent", newFeeChargePercent)
}

// ModifyFeeChargePercent is a paid mutator transaction binding the contract method 0xd38a85ad.
//
// Solidity: function modifyFeeChargePercent(uint24 newFeeChargePercent) returns()
func (_Swap *SwapSession) ModifyFeeChargePercent(newFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.ModifyFeeChargePercent(&_Swap.TransactOpts, newFeeChargePercent)
}

// ModifyFeeChargePercent is a paid mutator transaction binding the contract method 0xd38a85ad.
//
// Solidity: function modifyFeeChargePercent(uint24 newFeeChargePercent) returns()
func (_Swap *SwapTransactorSession) ModifyFeeChargePercent(newFeeChargePercent *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.ModifyFeeChargePercent(&_Swap.TransactOpts, newFeeChargePercent)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactor) SwapX2Y(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapX2Y", recipient, amount, lowPt, data)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapSession) SwapX2Y(recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapX2Y(&_Swap.TransactOpts, recipient, amount, lowPt, data)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x857f812f.
//
// Solidity: function swapX2Y(address recipient, uint128 amount, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactorSession) SwapX2Y(recipient common.Address, amount *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapX2Y(&_Swap.TransactOpts, recipient, amount, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactor) SwapX2YDesireY(opts *bind.TransactOpts, recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapX2YDesireY", recipient, desireY, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapSession) SwapX2YDesireY(recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapX2YDesireY(&_Swap.TransactOpts, recipient, desireY, lowPt, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0x59dd1436.
//
// Solidity: function swapX2YDesireY(address recipient, uint128 desireY, int24 lowPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactorSession) SwapX2YDesireY(recipient common.Address, desireY *big.Int, lowPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapX2YDesireY(&_Swap.TransactOpts, recipient, desireY, lowPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactor) SwapY2X(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapY2X", recipient, amount, highPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapSession) SwapY2X(recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapY2X(&_Swap.TransactOpts, recipient, amount, highPt, data)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x2c481252.
//
// Solidity: function swapY2X(address recipient, uint128 amount, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactorSession) SwapY2X(recipient common.Address, amount *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapY2X(&_Swap.TransactOpts, recipient, amount, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactor) SwapY2XDesireX(opts *bind.TransactOpts, recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapY2XDesireX", recipient, desireX, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapSession) SwapY2XDesireX(recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapY2XDesireX(&_Swap.TransactOpts, recipient, desireX, highPt, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0xf094685a.
//
// Solidity: function swapY2XDesireX(address recipient, uint128 desireX, int24 highPt, bytes data) returns(uint256 amountX, uint256 amountY)
func (_Swap *SwapTransactorSession) SwapY2XDesireX(recipient common.Address, desireX *big.Int, highPt *big.Int, data []byte) (*types.Transaction, error) {
	return _Swap.Contract.SwapY2XDesireX(&_Swap.TransactOpts, recipient, desireX, highPt, data)
}

// SwapAddLimitOrderIterator is returned from FilterAddLimitOrder and is used to iterate over the raw logs and unpacked data for AddLimitOrder events raised by the Swap contract.
type SwapAddLimitOrderIterator struct {
	Event *SwapAddLimitOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapAddLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapAddLimitOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapAddLimitOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapAddLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapAddLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapAddLimitOrder represents a AddLimitOrder event raised by the Swap contract.
type SwapAddLimitOrder struct {
	Owner         common.Address
	AddAmount     *big.Int
	AcquireAmount *big.Int
	Point         *big.Int
	ClaimSold     *big.Int
	ClaimEarn     *big.Int
	SellXEarnY    bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddLimitOrder is a free log retrieval operation binding the contract event 0x4f4658280ee6d0e8f09b5e436dacaca69ec5dd7c2ba05fb010d5145a3567cdad.
//
// Solidity: event AddLimitOrder(address indexed owner, uint128 addAmount, uint128 acquireAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) FilterAddLimitOrder(opts *bind.FilterOpts, owner []common.Address, point []*big.Int) (*SwapAddLimitOrderIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "AddLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return &SwapAddLimitOrderIterator{contract: _Swap.contract, event: "AddLimitOrder", logs: logs, sub: sub}, nil
}

// WatchAddLimitOrder is a free log subscription operation binding the contract event 0x4f4658280ee6d0e8f09b5e436dacaca69ec5dd7c2ba05fb010d5145a3567cdad.
//
// Solidity: event AddLimitOrder(address indexed owner, uint128 addAmount, uint128 acquireAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) WatchAddLimitOrder(opts *bind.WatchOpts, sink chan<- *SwapAddLimitOrder, owner []common.Address, point []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "AddLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapAddLimitOrder)
				if err := _Swap.contract.UnpackLog(event, "AddLimitOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddLimitOrder is a log parse operation binding the contract event 0x4f4658280ee6d0e8f09b5e436dacaca69ec5dd7c2ba05fb010d5145a3567cdad.
//
// Solidity: event AddLimitOrder(address indexed owner, uint128 addAmount, uint128 acquireAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) ParseAddLimitOrder(log types.Log) (*SwapAddLimitOrder, error) {
	event := new(SwapAddLimitOrder)
	if err := _Swap.contract.UnpackLog(event, "AddLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Swap contract.
type SwapBurnIterator struct {
	Event *SwapBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBurn represents a Burn event raised by the Swap contract.
type SwapBurn struct {
	Owner      common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	Liquidity  *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*SwapBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Burn", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &SwapBurnIterator{contract: _Swap.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *SwapBurn, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Burn", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBurn)
				if err := _Swap.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) ParseBurn(log types.Log) (*SwapBurn, error) {
	event := new(SwapBurn)
	if err := _Swap.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapCollectLimitOrderIterator is returned from FilterCollectLimitOrder and is used to iterate over the raw logs and unpacked data for CollectLimitOrder events raised by the Swap contract.
type SwapCollectLimitOrderIterator struct {
	Event *SwapCollectLimitOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapCollectLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapCollectLimitOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapCollectLimitOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapCollectLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapCollectLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapCollectLimitOrder represents a CollectLimitOrder event raised by the Swap contract.
type SwapCollectLimitOrder struct {
	Owner       common.Address
	Recipient   common.Address
	Point       *big.Int
	CollectDec  *big.Int
	CollectEarn *big.Int
	SellXEarnY  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCollectLimitOrder is a free log retrieval operation binding the contract event 0x7d3d0e34c86e56b4dcd993c09bbbf1b04527ab27b4365dffca10e0ded914e071.
//
// Solidity: event CollectLimitOrder(address indexed owner, address recipient, int24 indexed point, uint128 collectDec, uint128 collectEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) FilterCollectLimitOrder(opts *bind.FilterOpts, owner []common.Address, point []*big.Int) (*SwapCollectLimitOrderIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "CollectLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return &SwapCollectLimitOrderIterator{contract: _Swap.contract, event: "CollectLimitOrder", logs: logs, sub: sub}, nil
}

// WatchCollectLimitOrder is a free log subscription operation binding the contract event 0x7d3d0e34c86e56b4dcd993c09bbbf1b04527ab27b4365dffca10e0ded914e071.
//
// Solidity: event CollectLimitOrder(address indexed owner, address recipient, int24 indexed point, uint128 collectDec, uint128 collectEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) WatchCollectLimitOrder(opts *bind.WatchOpts, sink chan<- *SwapCollectLimitOrder, owner []common.Address, point []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "CollectLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapCollectLimitOrder)
				if err := _Swap.contract.UnpackLog(event, "CollectLimitOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectLimitOrder is a log parse operation binding the contract event 0x7d3d0e34c86e56b4dcd993c09bbbf1b04527ab27b4365dffca10e0ded914e071.
//
// Solidity: event CollectLimitOrder(address indexed owner, address recipient, int24 indexed point, uint128 collectDec, uint128 collectEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) ParseCollectLimitOrder(log types.Log) (*SwapCollectLimitOrder, error) {
	event := new(SwapCollectLimitOrder)
	if err := _Swap.contract.UnpackLog(event, "CollectLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapCollectLiquidityIterator is returned from FilterCollectLiquidity and is used to iterate over the raw logs and unpacked data for CollectLiquidity events raised by the Swap contract.
type SwapCollectLiquidityIterator struct {
	Event *SwapCollectLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapCollectLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapCollectLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapCollectLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapCollectLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapCollectLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapCollectLiquidity represents a CollectLiquidity event raised by the Swap contract.
type SwapCollectLiquidity struct {
	Owner      common.Address
	Recipient  common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectLiquidity is a free log retrieval operation binding the contract event 0xf69135213cd78fa4cffb855edf80272133f69bd8a6fb3236340a69b4d6e248e3.
//
// Solidity: event CollectLiquidity(address indexed owner, address recipient, int24 indexed leftPoint, int24 indexed rightPoint, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) FilterCollectLiquidity(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*SwapCollectLiquidityIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "CollectLiquidity", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &SwapCollectLiquidityIterator{contract: _Swap.contract, event: "CollectLiquidity", logs: logs, sub: sub}, nil
}

// WatchCollectLiquidity is a free log subscription operation binding the contract event 0xf69135213cd78fa4cffb855edf80272133f69bd8a6fb3236340a69b4d6e248e3.
//
// Solidity: event CollectLiquidity(address indexed owner, address recipient, int24 indexed leftPoint, int24 indexed rightPoint, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) WatchCollectLiquidity(opts *bind.WatchOpts, sink chan<- *SwapCollectLiquidity, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "CollectLiquidity", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapCollectLiquidity)
				if err := _Swap.contract.UnpackLog(event, "CollectLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectLiquidity is a log parse operation binding the contract event 0xf69135213cd78fa4cffb855edf80272133f69bd8a6fb3236340a69b4d6e248e3.
//
// Solidity: event CollectLiquidity(address indexed owner, address recipient, int24 indexed leftPoint, int24 indexed rightPoint, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) ParseCollectLiquidity(log types.Log) (*SwapCollectLiquidity, error) {
	event := new(SwapCollectLiquidity)
	if err := _Swap.contract.UnpackLog(event, "CollectLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapDecLimitOrderIterator is returned from FilterDecLimitOrder and is used to iterate over the raw logs and unpacked data for DecLimitOrder events raised by the Swap contract.
type SwapDecLimitOrderIterator struct {
	Event *SwapDecLimitOrder // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapDecLimitOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapDecLimitOrder)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapDecLimitOrder)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapDecLimitOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapDecLimitOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapDecLimitOrder represents a DecLimitOrder event raised by the Swap contract.
type SwapDecLimitOrder struct {
	Owner          common.Address
	DecreaseAmount *big.Int
	Point          *big.Int
	ClaimSold      *big.Int
	ClaimEarn      *big.Int
	SellXEarnY     bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDecLimitOrder is a free log retrieval operation binding the contract event 0x3736ba81d13006f6ea2012ba3e287f087169b55d90a9defb5966fe9eb830d7ea.
//
// Solidity: event DecLimitOrder(address indexed owner, uint128 decreaseAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) FilterDecLimitOrder(opts *bind.FilterOpts, owner []common.Address, point []*big.Int) (*SwapDecLimitOrderIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "DecLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return &SwapDecLimitOrderIterator{contract: _Swap.contract, event: "DecLimitOrder", logs: logs, sub: sub}, nil
}

// WatchDecLimitOrder is a free log subscription operation binding the contract event 0x3736ba81d13006f6ea2012ba3e287f087169b55d90a9defb5966fe9eb830d7ea.
//
// Solidity: event DecLimitOrder(address indexed owner, uint128 decreaseAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) WatchDecLimitOrder(opts *bind.WatchOpts, sink chan<- *SwapDecLimitOrder, owner []common.Address, point []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var pointRule []interface{}
	for _, pointItem := range point {
		pointRule = append(pointRule, pointItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "DecLimitOrder", ownerRule, pointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapDecLimitOrder)
				if err := _Swap.contract.UnpackLog(event, "DecLimitOrder", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecLimitOrder is a log parse operation binding the contract event 0x3736ba81d13006f6ea2012ba3e287f087169b55d90a9defb5966fe9eb830d7ea.
//
// Solidity: event DecLimitOrder(address indexed owner, uint128 decreaseAmount, int24 indexed point, uint128 claimSold, uint128 claimEarn, bool sellXEarnY)
func (_Swap *SwapFilterer) ParseDecLimitOrder(log types.Log) (*SwapDecLimitOrder, error) {
	event := new(SwapDecLimitOrder)
	if err := _Swap.contract.UnpackLog(event, "DecLimitOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the Swap contract.
type SwapFlashIterator struct {
	Event *SwapFlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlash represents a Flash event raised by the Swap contract.
type SwapFlash struct {
	Sender    common.Address
	Recipient common.Address
	AmountX   *big.Int
	AmountY   *big.Int
	PaidX     *big.Int
	PaidY     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_Swap *SwapFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*SwapFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashIterator{contract: _Swap.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_Swap *SwapFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *SwapFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlash)
				if err := _Swap.contract.UnpackLog(event, "Flash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amountX, uint256 amountY, uint256 paidX, uint256 paidY)
func (_Swap *SwapFilterer) ParseFlash(log types.Log) (*SwapFlash, error) {
	event := new(SwapFlash)
	if err := _Swap.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Swap contract.
type SwapMintIterator struct {
	Event *SwapMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapMint represents a Mint event raised by the Swap contract.
type SwapMint struct {
	Sender     common.Address
	Owner      common.Address
	LeftPoint  *big.Int
	RightPoint *big.Int
	Liquidity  *big.Int
	AmountX    *big.Int
	AmountY    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (*SwapMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Mint", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return &SwapMintIterator{contract: _Swap.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SwapMint, owner []common.Address, leftPoint []*big.Int, rightPoint []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var leftPointRule []interface{}
	for _, leftPointItem := range leftPoint {
		leftPointRule = append(leftPointRule, leftPointItem)
	}
	var rightPointRule []interface{}
	for _, rightPointItem := range rightPoint {
		rightPointRule = append(rightPointRule, rightPointItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Mint", ownerRule, leftPointRule, rightPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapMint)
				if err := _Swap.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed leftPoint, int24 indexed rightPoint, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Swap *SwapFilterer) ParseMint(log types.Log) (*SwapMint, error) {
	event := new(SwapMint)
	if err := _Swap.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Swap contract.
type SwapSwapIterator struct {
	Event *SwapSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapSwap represents a Swap event raised by the Swap contract.
type SwapSwap struct {
	TokenX       common.Address
	TokenY       common.Address
	Fee          *big.Int
	SellXEarnY   bool
	AmountX      *big.Int
	AmountY      *big.Int
	CurrentPoint *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x0fe977d619f8172f7fdbe8bb8928ef80952817d96936509f67d66346bc4cd10f.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY, int24 currentPoint)
func (_Swap *SwapFilterer) FilterSwap(opts *bind.FilterOpts, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (*SwapSwapIterator, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Swap", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &SwapSwapIterator{contract: _Swap.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0fe977d619f8172f7fdbe8bb8928ef80952817d96936509f67d66346bc4cd10f.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY, int24 currentPoint)
func (_Swap *SwapFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapSwap, tokenX []common.Address, tokenY []common.Address, fee []*big.Int) (event.Subscription, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Swap", tokenXRule, tokenYRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapSwap)
				if err := _Swap.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0x0fe977d619f8172f7fdbe8bb8928ef80952817d96936509f67d66346bc4cd10f.
//
// Solidity: event Swap(address indexed tokenX, address indexed tokenY, uint24 indexed fee, bool sellXEarnY, uint256 amountX, uint256 amountY, int24 currentPoint)
func (_Swap *SwapFilterer) ParseSwap(log types.Log) (*SwapSwap, error) {
	event := new(SwapSwap)
	if err := _Swap.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
