package mock

import (
	"github.com/a1146910248/mixchain/mvm/common"
	"github.com/a1146910248/mixchain/mvm/types"
	"math/big"
)

func GetHeader(blockNum, difficulty int64, timeStamp uint64) *types.Header {
	return &types.Header{
		ParentHash:       common.Hash{},
		UncleHash:        common.Hash{},
		Coinbase:         common.Address{},
		Root:             common.Hash{},
		TxHash:           common.Hash{},
		ReceiptHash:      common.Hash{},
		Bloom:            types.Bloom{},
		Difficulty:       new(big.Int).Set(big.NewInt(difficulty)),
		Number:           new(big.Int).Set(big.NewInt(blockNum)),
		GasLimit:         0xfffffffffffffff,
		GasUsed:          0,
		Time:             timeStamp,
		Extra:            nil,
		MixDigest:        common.Hash{},
		Nonce:            types.BlockNonce{},
		BaseFee:          nil,
		WithdrawalsHash:  nil,
		BlobGasUsed:      nil,
		ExcessBlobGas:    nil,
		ParentBeaconRoot: nil,
	}
}
