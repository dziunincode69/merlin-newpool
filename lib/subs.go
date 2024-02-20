package lib

import (
	"context"
	"fmt"
	"log"
	"merlin-newpool/client"
	"merlin-newpool/utils"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

func EventSubscribe(logchan chan types.Log) []ethereum.Subscription {
	// var
	var subs []ethereum.Subscription
	c, err := client.RPC().SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Topics: [][]common.Hash{{common.HexToHash("0xf04da67755adf58739649e2fb9949a6328518141b7ac9e44aa10320688b04900")}},
	}, logchan)
	if err != nil {
		log.Fatal(err)
	}
	d, err := client.RPC().SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Topics: [][]common.Hash{{common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"), common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")}},
	}, logchan)
	if err != nil {
		log.Fatal(err)
	}
	d2, err := client.RPC().SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Topics: [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")}},
	}, logchan)
	if err != nil {
		log.Fatal(err)
	}
	subs = append(subs, c, d, d2)
	return subs
}
func Subscribe(poolinit *Caller) bool {
	logchan := make(chan types.Log)
	triggerchan := make(chan bool)
	EventSubscribe(logchan)
	// go func() {
	// 	cronjob.AddFunc("@every 10m", func() {
	// 		log.Println("Restarting connection...")
	// 		for _, v := range subs {
	// 			wg.Add(1)
	// 			go func(s ethereum.Subscription) {
	// 				defer wg.Done()
	// 				s.Unsubscribe()
	// 			}(v)
	// 		}
	// 		wg.Wait()
	// 		triggerchan <- true
	// 		cronjob.Stop()
	// 	})
	// 	cronjob.Start()
	// }()

	go func(logchan chan types.Log) {
		for vLog := range logchan {
			if vLog.Topics[0].Hex() == "0xf04da67755adf58739649e2fb9949a6328518141b7ac9e44aa10320688b04900" {
				trx, ispending, err := client.RPC().TransactionByHash(context.Background(), vLog.TxHash)
				if err != nil {
					log.Fatal(err)

				}
				log.Println("Found New Event", vLog.TxHash.String(), "Pending: ", ispending, "Block: ", vLog.BlockNumber, "Data: ", hexutil.Encode(trx.Data()))
				resultdecode, err := Decode(hexutil.Encode(trx.Data()))
				if err != nil {
					log.Fatal(err)

				}
				pooladdress, err := poolinit.GetPoolAddress(&Metas{
					TokenX: resultdecode.TokenX,
					TokenY: resultdecode.TokenY,
					Fee:    resultdecode.Fee,
				})
				log.Println("Got Pool Address: ", pooladdress.String(), "From Token: ", resultdecode.TokenX.String())
				if err != nil {
					log.Fatal(err)

				}
				newerc, err := InitERC(resultdecode.TokenX)
				if err != nil {
					log.Fatal(err)

				}
				ercres, err := newerc.GetTokenDetail()
				if err != nil {
					log.Fatal(err)
				}
				msg := "*🔮 MERLIN NEW CREATE POOL 🔮*\n\n" +
					"*Token Name:* " + utils.EscapeMarkdownV2(ercres.Name) + "\n" +
					"*Token Symbol:* " + utils.EscapeMarkdownV2(ercres.Symbol) + "\n" +
					"*Total Supply:* " + utils.EscapeMarkdownV2(ercres.TotalSupply.String()) + "\n" +
					"*Decimals:* " + utils.EscapeMarkdownV2(strconv.Itoa(int(ercres.Dec))) + "\n\n" +
					"*Token Address:* `" + utils.EscapeMarkdownV2(resultdecode.TokenX.Hex()) + "`\n" +
					"*Owner:* `" + utils.EscapeMarkdownV2(ercres.Owner.Hex()) + "`\n" +
					"*Pool Address:* `" + utils.EscapeMarkdownV2(pooladdress.Hex()) + "`\n" +
					"*Fee:* `" + fmt.Sprintf("%d", resultdecode.Fee.Int64()) + "`\n\n" +
					"[Owner](https://scan.merlinchain.io/address/" + ercres.Owner.Hex() + ")  " +
					"[Token Address](https://scan.merlinchain.io/token/" + resultdecode.TokenX.Hex() + ")  " +
					"[Pool Address](https://scan.merlinchain.io/address/" + pooladdress.Hex() + ")"
				fmt.Println("TokenAddress: ", resultdecode.TokenX.String(), "Name: ", ercres.Name, "Symbol: ", ercres.Symbol, "Owner: ", ercres.Owner.String(), "PoolAddress: ", pooladdress.String(), "Fee: ", resultdecode.Fee.Int64())
				SendLog(msg, "@merlinnewpool")
			} else if vLog.Topics[0].Hex() == "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0" || vLog.Topics[0].Hex() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
				trx, err := client.RPC().TransactionReceipt(context.Background(), vLog.TxHash)
				if err != nil {
					log.Fatal(err)
				}
				if CheckTKNInserted(trx.ContractAddress) {
					if trx.ContractAddress.String() != "0x0000000000000000000000000000000000000000" {
						log.Println("Found New Token", vLog.TxHash.String(), "Token Address: ", common.HexToAddress(trx.ContractAddress.String()).String())
						newerc, err := InitERC(trx.ContractAddress)
						if err != nil {
							triggerchan <- false
						} else {
							ercres, err := newerc.GetTokenDetail()
							if err != nil {
								triggerchan <- false
							} else {
								msg := "*🔮 MERLIN NEW TOKEN CREATED 🔮*\n\n" +
									"*Token Name:* " + utils.EscapeMarkdownV2(ercres.Name) + "\n" +
									"*Token Symbol:* " + utils.EscapeMarkdownV2(ercres.Symbol) + "\n" +
									"*Token Address:* `" + utils.EscapeMarkdownV2(trx.ContractAddress.Hex()) + "`\n" +
									"*Total Supply:* " + utils.EscapeMarkdownV2(ercres.TotalSupply.String()) + "\n" +
									"*Decimals:* " + utils.EscapeMarkdownV2(strconv.Itoa(int(ercres.Dec))) + "\n" +
									"*Owner:* `" + utils.EscapeMarkdownV2(ercres.Owner.Hex()) + "`\n\n" +
									"[Token Address](https://scan.merlinchain.io/token/" + trx.ContractAddress.Hex() + ")  " +
									"[Owner](https://scan.merlinchain.io/address/" + ercres.Owner.Hex() + ")  " +
									"[TX HASH](https://scan.merlinchain.io/tx/" + vLog.TxHash.Hex() + ")  "
								fmt.Println("TokenAddress: ", trx.ContractAddress.String(), "Name: ", ercres.Name, "Symbol: ", ercres.Symbol, "Owner: ", ercres.Owner.String())
								InsertNewTkn(trx.ContractAddress.Hex(), ercres.Name, ercres.Symbol, ercres.Owner.String())
								// InsertTknLog(resultdecode.TokenX.String(), ercres.Name, ercres.Symbol, ercres.Owner.String(), pooladdress.String(), resultdecode.Fee.Int64(), 0)
								SendLog(msg, "@merlinnewtoken")
							}
						}
					}
				}

			}
		}
		triggerchan <- true
	}(logchan)

	return <-triggerchan
}
