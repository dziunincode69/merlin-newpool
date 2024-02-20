package lib

import (
	"math/big"
	"merlin-newpool/client"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeta-chain/protocol-contracts/pkg/openzeppelin/contracts/token/erc20/erc20.sol"
)

type TokenDetail struct {
	Name        string
	Symbol      string
	Dec         uint8
	TotalSupply *big.Int
	Owner       common.Address
}

type request struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

type CallerToken struct {
	Caller       *erc20.ERC20
	TokenAddress common.Address
}

func GetOwner(token string) string {
	var getOwner string
	req := request{token, "0x893d20e8"}
	if err := client.RPCCUSTOM().Call(&getOwner, "eth_call", req, "latest"); err != nil || getOwner == "0x0000000000000000000000000000000000000000" {
		if err = client.RPCCUSTOM().Call(&getOwner, "eth_call", request{token, "0xb4a99a4e"}, "latest"); err != nil || getOwner == "0x0000000000000000000000000000000000000000" {
			if err = client.RPCCUSTOM().Call(&getOwner, "eth_call", request{token, "0x8da5cb5b"}, "latest"); err != nil || getOwner == "0x0000000000000000000000000000000000000000" {
				return "0x0000000000000000000000000000000000000000"
			}
		}
	}
	getOwner = strings.ToLower(common.HexToAddress(getOwner).Hex())
	return getOwner
}

func InitERC(token common.Address) (*CallerToken, error) {
	erc20, err := erc20.NewERC20(token, client.RPC())
	if err != nil {
		return nil, err
	}
	return &CallerToken{
		Caller:       erc20,
		TokenAddress: token,
	}, nil

}
func (C *CallerToken) GetTokenDetail() (*TokenDetail, error) {
	name, err := C.Caller.Name(nil)
	if err != nil {
		return nil, err
	}
	symbol, err := C.Caller.Symbol(nil)
	if err != nil {
		return nil, err
	}
	dec, err := C.Caller.Decimals(nil)
	if err != nil {
		return nil, err
	}
	totalSupply, err := C.Caller.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	owner := GetOwner(C.TokenAddress.String())
	return &TokenDetail{
		Name:        name,
		Symbol:      symbol,
		Dec:         dec,
		TotalSupply: totalSupply,
		Owner:       common.HexToAddress(owner),
	}, nil

}
