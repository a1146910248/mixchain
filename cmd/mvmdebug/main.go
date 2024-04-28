package main

import (
	"encoding/hex"
	"fmt"
	"github.com/a1146910248/mixchain/mvm"
	"github.com/a1146910248/mixchain/mvm/abi"
	"github.com/a1146910248/mixchain/mvm/mock"
	"github.com/a1146910248/mixchain/mvm/params"
	"github.com/a1146910248/mixchain/mvm/state"
	"github.com/a1146910248/mixchain/mvm/vm"
	"github.com/holiman/uint256"
	"math/big"
	"reflect"
	"strings"
)

// 2e64cec1 : retrieve, 6057361d000000000000000000000000000000000000000000000000000000000000007b: store 123
var input, _ = hex.DecodeString("2e64cec1")

func main() {
	// return
	//updateContract()
	// 创建账户State
	stateDb, err := state.TryLoadFromDisk()
	if err != nil {
		panic(err)
	}
	blockCtx := mvm.NewEVMBlockContext(mock.GetHeader(100, 1, 1200000))
	txCtx := mvm.NewEVMTxContext(mock.GetMessage(normalAccount))
	vmenv := vm.NewEVM(blockCtx, txCtx, stateDb, params.AllEthashProtocolChanges, vm.Config{})

	//ret, caddr, leftgas, err := vmenv.Create(vm.AccountRef(normalAccount), helloCode, 1000000, new(uint256.Int))
	//fmt.Printf("usedGas: %v,caddr:%v, err: %v, len(ret): %v \n", 1000000-leftgas, caddr, err, len(ret))
	ret, leftgas, err := vmenv.Call(vm.AccountRef(normalAccount), helloWorldcontactAccont, input, 1000000, new(uint256.Int))
	//appen := make([]byte, 16)
	//ret = append(ret, appen...)
	fmt.Printf("usedGas: %v, err: %v, len(ret): %v \n", 1000000-leftgas, err, len(ret))
	//fmt.Printf("ret: %v, usedGas: %v, err: %v, len(ret): %v, hexret: %v, ", ret, 1000000-leftgas, err, len(ret), hex.EncodeToString(ret))
	abiObjet, _ := abi.JSON(strings.NewReader(storeContractABIJson))

	// begin, length, _ := lengthPrefixPointsTo(0, ret)

	value := big.NewInt(0) //new(*big.Int)
	fmt.Println(abiObjet.UnpackIntoInterface(&value, "retrieve", ret))
	//fmt.Println(unpackAtomic(&restult, string(ret[begin:begin+length])))
	println(value.String())
	fmt.Println(stateDb.Commit())
}

func lengthPrefixPointsTo(index int, output []byte) (start int, length int, err error) {
	bigOffsetEnd := big.NewInt(0).SetBytes(output[index : index+32])
	bigOffsetEnd.Add(bigOffsetEnd, big.NewInt(32))
	outputLength := big.NewInt(int64(len(output)))

	if bigOffsetEnd.Cmp(outputLength) > 0 {
		return 0, 0, fmt.Errorf("abi: cannot marshal in to go slice: offset %v would go over slice boundary (len=%v)", bigOffsetEnd, outputLength)
	}

	if bigOffsetEnd.BitLen() > 63 {
		return 0, 0, fmt.Errorf("abi offset larger than int64: %v", bigOffsetEnd)
	}

	offsetEnd := int(bigOffsetEnd.Uint64())
	lengthBig := big.NewInt(0).SetBytes(output[offsetEnd-32 : offsetEnd])

	totalSize := big.NewInt(0)
	totalSize.Add(totalSize, bigOffsetEnd)
	totalSize.Add(totalSize, lengthBig)
	if totalSize.BitLen() > 63 {
		return 0, 0, fmt.Errorf("abi length larger than int64: %v", totalSize)
	}

	if totalSize.Cmp(outputLength) > 0 {
		return 0, 0, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %v require %v", outputLength, totalSize)
	}
	start = int(bigOffsetEnd.Uint64())
	length = int(lengthBig.Uint64())
	return
}

func unpackAtomic(v interface{}, marshalledValues interface{}) error {

	elem := reflect.ValueOf(v).Elem()
	// kind := elem.Kind()
	reflectValue := reflect.ValueOf(marshalledValues)
	return set(elem, reflectValue)
}

func set(dst, src reflect.Value) error {
	dstType := dst.Type()
	srcType := src.Type()
	switch {
	case dstType.AssignableTo(srcType):
		dst.Set(src)
	case dstType.Kind() == reflect.Interface:
		dst.Set(src)
	case dstType.Kind() == reflect.Ptr:
		return set(dst.Elem(), src)
	default:
		return fmt.Errorf("abi: cannot unmarshal %v in to %v", src.Type(), dst.Type())
	}
	return nil
}
