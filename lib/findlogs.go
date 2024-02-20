package lib

import (
	"context"
	"fmt"
	"log"
	"merlin-newpool/client"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func GetCreatePoolFromLogsByOwnerAddress(devwallet string) {
	query := ethereum.FilterQuery{
		// FromBlock: big.NewInt(100809),
		Topics: [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"), common.HexToHash("0x000000000000000000000000df086fb1be147b3052455d553189d3b585da984a")}},
	}
	res, err := client.RPC().FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		return
	}
	// fmt.Println(res)
	for _, v := range res {
		fmt.Println("Found Recent Pool Creation")
		fmt.Println(v.TxHash)
		// fmt.Println(v)
		// fmt.Println(v.Topics[2].Hex())
		// fmt.Println(v.Topics[1].Hex())
		// fmt.Println(v.Topics[0].Hex())
		// fmt.Println(v.BlockNumber)
		// fmt.Println(v.TxHash.Hex())
		// fmt.Println(v.BlockHash.Hex())
		// fmt.Println(v.Address.Hex())
		// fmt.Println(v.Data)
		// fmt.Println(v.Removed)
	}
}
