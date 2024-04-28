package mock

import (
	"github.com/a1146910248/mixchain/mvm"
	"github.com/a1146910248/mixchain/mvm/common"
	"math/big"
)

func GetMessage(from common.Address) *mvm.Message {
	return &mvm.Message{
		To:                nil,
		From:              from,
		Nonce:             0,
		Value:             nil,
		GasLimit:          0,
		GasPrice:          new(big.Int).Set(big.NewInt(10)),
		GasFeeCap:         nil,
		GasTipCap:         nil,
		Data:              nil,
		AccessList:        nil,
		BlobGasFeeCap:     nil,
		BlobHashes:        nil,
		SkipAccountChecks: false,
	}
}
