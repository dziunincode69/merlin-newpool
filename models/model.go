package models

import "github.com/kamva/mgm/v3"

type LogToken struct {
	mgm.DefaultModel `bson:",inline"`
	TokenAddress     string `json:"tokenAddress" bson:"tokenAddress"`
	Fee              int64  `json:"fee" bson:"fee"`
	Name             string `json:"name" bson:"name"`
	Symbol           string `json:"symbol" bson:"symbol"`
	PoolIds          int64  `json:"poolId" bson:"poolId"`
	Owner            string `json:"owner" bson:"owner"`
	PoolAddress      string `json:"poolAddress" bson:"poolAddress"`
}
type NewToken struct {
	mgm.DefaultModel `bson:",inline"`
	TokenAddress     string `json:"tokenAddress" bson:"tokenAddress"`
	Name             string `json:"name" bson:"name"`
	Symbol           string `json:"symbol" bson:"symbol"`
	Owner            string `json:"owner" bson:"owner"`
}

func NewModelLogTkn(tknAddress, name, symbol, Owner, PoolAddress string, Fee, Poolid int64) *LogToken {
	return &LogToken{
		TokenAddress: tknAddress,
		Fee:          Fee,
		Name:         name,
		Symbol:       symbol,
		PoolIds:      Poolid,
		Owner:        Owner,
		PoolAddress:  PoolAddress,
	}
}

func NewModelNewToken(tknAddress, name, symbol, Owner string) *LogToken {
	return &LogToken{
		TokenAddress: tknAddress,
		Name:         name,
		Symbol:       symbol,
		Owner:        Owner,
	}
}
