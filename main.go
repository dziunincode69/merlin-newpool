package main

import (
	"fmt"
	"log"
	"math/big"
	"merlin-newpool/client"
	"merlin-newpool/lib"
	"merlin-newpool/models"
	"merlin-newpool/utils"
	"strconv"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	client.ConnectRpc("https://rpc.merlinchain.io/")
	client.ConnectRPCCustom("https://rpc.merlinchain.io/")

	client.ConnectMongo()

}

//mongo merlin Saninkicker123

func main() {
	var logtkn models.LogToken
	err := mgm.Coll(&logtkn).FindOne(mgm.Ctx(), bson.D{}, options.FindOne().SetSort(bson.D{{"created_at", -1}})).Decode(&logtkn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(logtkn)
	poolinit, err := lib.BindNewContract()
	if err != nil {
		log.Fatal(err)
	}
	start := logtkn.PoolIds
	count := 0
	for {
		res, err := poolinit.GetPoolMetas(big.NewInt(int64(start)))
		if err != nil {
			log.Fatal(err)
		}
		pooladdress, err := poolinit.GetPoolAddress(res)
		if err != nil {
			log.Fatal(err, "error GetPoolAddress")
		}
		count++
		if res.TokenX.String() != "0x0000000000000000000000000000000000000000" {
			count = 0
			newerc, err := lib.InitERC(res.TokenX)
			if err != nil {
				log.Fatal(err)
			}
			ercres, err := newerc.GetTokenDetail()
			if err != nil {
				log.Fatal(err, "error GetTokenDetail")
			}
			msg := "*🔮 MERLIN NEW POOL 🔮*\n\n" +
				"*Token Name:* *" + utils.EscapeMarkdownV2(ercres.Name) + "*\n" +
				"*Token Symbol:* *" + utils.EscapeMarkdownV2(ercres.Symbol) + "*\n" +
				"*Total Supply:* " + utils.EscapeMarkdownV2(ercres.TotalSupply.String()) + "\n" +
				"*Decimals:* " + utils.EscapeMarkdownV2(strconv.Itoa(int(ercres.Dec))) + "\n\n" +
				"*Token Address:* [" + utils.EscapeMarkdownV2(res.TokenX.Hex()) + "](https://scan.merlinchain.io/token/" + res.TokenX.Hex() + ")\n" +
				"*Owner:* `" + utils.EscapeMarkdownV2(ercres.Owner.Hex()) + "`\n" +
				"*Pool Address:* `" + utils.EscapeMarkdownV2(pooladdress.Hex()) + "`\n" +
				"*PoolFee:* `" + fmt.Sprintf("%d", res.Fee.Int64()) + "`\n"

			// Send the message to Telegram
			fmt.Println("TokenAddress: ", res.TokenX.String(), "Name: ", ercres.Name, "Symbol: ", ercres.Symbol, "Owner: ", ercres.Owner.String(), "PoolAddress: ", pooladdress.String(), "Fee: ", res.Fee.Int64(), "PoolId: ", int64(start))
			lib.InsertTknLog(res.TokenX.String(), ercres.Name, ercres.Symbol, ercres.Owner.String(), pooladdress.String(), res.Fee.Int64(), int64(start))
			lib.SendLog(msg, "@merlinnewpool")
			// lib.SendLog(msg, "1644665024")
			start++
		}
		time.Sleep(2 * time.Second)
		fmt.Print("\r", res.TokenX, " Count:", count, start)
	}

}
