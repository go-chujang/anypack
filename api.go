package anypack

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func Pack(abi abi.ABI, name string, input any) (packed []byte, err error) {
	method, exist := abi.Methods[name]
	if !exist {
		return nil, errors.New("method not found")
	}

	args, err := toArgs(method.Inputs, input)
	if err != nil {
		return nil, err
	}
	return abi.Pack(method.Name, args...)
}

func PackById(abi abi.ABI, id []byte, input any) (packed []byte, err error) {
	method, err := abi.MethodById(id)
	if err != nil {
		return nil, err
	}
	args, err := toArgs(method.Inputs, input)
	if err != nil {
		return nil, err
	}
	return abi.Pack(method.Name, args...)
}

func ToFunctionTy(id []byte, contractAddress common.Address) (funtionTy [24]byte, ok bool) {
	if len(id) != 4 {
		return [24]byte{}, false
	}
	functionTyBytes := append(contractAddress.Bytes(), id...)
	copy(funtionTy[:], functionTyBytes[0:24])
	return funtionTy, true
}
