package lib

import (
	"encoding/hex"
	"math/big"
	"merlin-newpool/interact"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type CreatePoolParams struct {
	TokenX       common.Address // Assuming path is a byte slice
	TokenY       common.Address // Ethereum address
	Fee          *big.Int       // Big integer for amount
	InitialPoint *big.Int       // Big integer for minimum acquired
}

func Decode(input string) (*CreatePoolParams, error) {

	abi, err := abi.JSON(strings.NewReader(interact.LiquiditymanagerABI))
	if err != nil {
		return nil, err

	}

	decodedSig, err := hex.DecodeString(input[2:10])
	if err != nil {
		return nil, err

	}

	method, err := abi.MethodById(decodedSig)
	if err != nil {
		return nil, err

	}
	decodedData, err := hex.DecodeString(input[10:])
	if err != nil {
		return nil, err
	}
	// hexCode := input[0:10]

	infoTx, _ := method.Inputs.Unpack(decodedData)
	// log.Println(hexCode, method.Name, infoTx)
	return &CreatePoolParams{
		TokenX:       common.HexToAddress(infoTx[0].(common.Address).Hex()),
		TokenY:       common.HexToAddress(infoTx[1].(common.Address).Hex()),
		Fee:          infoTx[2].(*big.Int),
		InitialPoint: infoTx[3].(*big.Int),
	}, nil
}
