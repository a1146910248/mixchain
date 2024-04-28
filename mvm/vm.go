package mvm

import (
	"github.com/a1146910248/mixchain/mvm/common"
	"github.com/a1146910248/mixchain/mvm/tracing"
	"github.com/a1146910248/mixchain/mvm/types"
	"github.com/a1146910248/mixchain/mvm/vm"

	"github.com/holiman/uint256"
	"math/big"
)

// A Message contains the data derived from a single transaction that is relevant to state
// processing.
type Message struct {
	To            *common.Address
	From          common.Address
	Nonce         uint64
	Value         *big.Int
	GasLimit      uint64
	GasPrice      *big.Int
	GasFeeCap     *big.Int
	GasTipCap     *big.Int
	Data          []byte
	AccessList    types.AccessList
	BlobGasFeeCap *big.Int
	BlobHashes    []common.Hash

	// When SkipAccountChecks is true, the message nonce is not checked against the
	// account nonce in state. It also disables checking that the sender is an EOA.
	// This field will be set to true for operations like RPC eth_call.
	SkipAccountChecks bool
}

// NewEVMBlockContext creates a new context for use in the EVM.
func NewEVMBlockContext(header *types.Header) vm.BlockContext {
	var (
		baseFee     *big.Int
		blobBaseFee *big.Int
		random      *common.Hash
	)

	if header.BaseFee != nil {
		baseFee = new(big.Int).Set(header.BaseFee)
	}
	if header.ExcessBlobGas != nil {
		blobBaseFee = new(big.Int)
	}
	if header.Difficulty.Sign() == 0 {
		random = &header.MixDigest
	}
	return vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     GetHashFn(),
		Coinbase:    header.Coinbase,
		BlockNumber: new(big.Int).Set(header.Number),
		Time:        header.Time,
		Difficulty:  new(big.Int).Set(header.Difficulty),
		BaseFee:     baseFee,
		BlobBaseFee: blobBaseFee,
		GasLimit:    header.GasLimit,
		Random:      random,
	}
}

// NewEVMTxContext creates a new transaction context for a single transaction.
func NewEVMTxContext(msg *Message) vm.TxContext {
	ctx := vm.TxContext{
		Origin:     msg.From,
		GasPrice:   new(big.Int).Set(msg.GasPrice),
		BlobHashes: msg.BlobHashes,
	}
	if msg.BlobGasFeeCap != nil {
		ctx.BlobFeeCap = new(big.Int).Set(msg.BlobGasFeeCap)
	}
	return ctx
}

// GetHashFn returns a GetHashFunc which retrieves header hashes by number
func GetHashFn() func(n uint64) common.Hash {
	// Cache will initially contain [refHash.parent],
	// Then fill up with [refHash.p, refHash.pp, refHash.ppp, ...]
	//var cache []common.Hash

	return func(n uint64) common.Hash {
		//if ref.Number.Uint64() <= n {
		//	// This situation can happen if we're doing tracing and using
		//	// block overrides.
		//	return common.Hash{}
		//}
		//// If there's no hash cache yet, make one
		//if len(cache) == 0 {
		//	cache = append(cache, ref.ParentHash)
		//}
		//if idx := ref.Number.Uint64() - n - 1; idx < uint64(len(cache)) {
		//	return cache[idx]
		//}
		//// No luck in the cache, but we can start iterating from the last element we already know
		//lastKnownHash := cache[len(cache)-1]
		//lastKnownNumber := ref.Number.Uint64() - uint64(len(cache))
		//
		//for {
		//	header := chain.GetHeader(lastKnownHash, lastKnownNumber)
		//	if header == nil {
		//		break
		//	}
		//	cache = append(cache, header.ParentHash)
		//	lastKnownHash = header.ParentHash
		//	lastKnownNumber = header.Number.Uint64() - 1
		//	if n == lastKnownNumber {
		//		return lastKnownHash
		//	}
		//}
		return common.Hash{}
	}
}

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db vm.StateDB, addr common.Address, amount *uint256.Int) bool {
	return db.GetBalance(addr).Cmp(amount) >= 0
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db vm.StateDB, sender, recipient common.Address, amount *uint256.Int) {
	db.SubBalance(sender, amount, tracing.BalanceChangeTransfer)
	db.AddBalance(recipient, amount, tracing.BalanceChangeTransfer)
}
