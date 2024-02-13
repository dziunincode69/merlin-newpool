package lib

import (
	"math/big"
	"merlin-newpool/client"
	"merlin-newpool/interact"
	"merlin-newpool/utils"

	"github.com/ethereum/go-ethereum/common"
)

type Caller struct {
	Liquiditymanager *interact.LiquiditymanagerCaller
}
type Metas struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}

func BindNewContract() (*Caller, error) {
	c, err := interact.NewLiquiditymanagerCaller(utils.LIQUIDITY_MANAGER, client.RPC())
	if err != nil {
		return nil, err
	}
	return &Caller{
		Liquiditymanager: c,
	}, nil
}
func (c Caller) GetPoolMetas(num *big.Int) (*Metas, error) {

	poolmetas, err := c.Liquiditymanager.PoolMetas(nil, num)
	if err != nil {
		return nil, err
	}
	return &Metas{
		TokenX: poolmetas.TokenX,
		TokenY: poolmetas.TokenY,
		Fee:    poolmetas.Fee,
	}, nil
}

func (c Caller) GetPoolAddress(meta *Metas) (common.Address, error) {
	pooladdress, err := c.Liquiditymanager.Pool(nil, meta.TokenX, meta.TokenY, meta.Fee)
	if err != nil {
		return common.Address{}, err
	}
	return pooladdress, nil
}
