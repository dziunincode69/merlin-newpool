package client

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	rpcClient  *ethclient.Client
	err        error
	rpcClient2 *rpc.Client
	From       common.Address
)

func ConnectMongo() {
	err = mgm.SetDefaultConfig(nil, "merlin-manesty", options.Client().ApplyURI("mongodb+srv://merlin:Saninkicker123@cluster0.5kwpu67.mongodb.net/")) //cloud@bowdown.fun
	if err != nil {
		fmt.Println("Fatal err MongoDB", err)
	}
}

func ConnectRpc(rpcUrl string) {
	rpcClient, err = ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
}
func ConnectRPCCustom(rpcUrl string) {
	rpcClient2, err = rpc.Dial(rpcUrl)

}
func RPCCUSTOM() *rpc.Client {
	return rpcClient2
}

func RPC() *ethclient.Client {
	return rpcClient
}
