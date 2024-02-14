package main

import (
	"fmt"
	"log"
	"merlin-newpool/client"
	"merlin-newpool/lib"
)

func init() {
	client.ConnectRpc("ws://142.132.192.47:8133/")
	client.ConnectRPCCustom("ws://142.132.192.47:8133/")
	client.ConnectMongo()
}

//mongo merlin Saninkicker123

func main() {
	poolinit, err := lib.BindNewContract()
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("bot Started Merlin New Pool & New Token")
		subs := lib.Subscribe(poolinit)
		if !subs {
			fmt.Printf("")
		}
	}

}
