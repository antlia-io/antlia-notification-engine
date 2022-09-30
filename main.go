package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antlia-io/antlia-notification-engine/config"
	"github.com/antlia-io/antlia-notification-engine/repository/mongodb"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	conf := config.LoadConfiguration("./config.json")

	client, err := ethclient.Dial(conf.NetworkWebSocket)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(conf.ContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	db := mongodb.SharedStore()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.TxHash, vLog.Data) // pointer to event log
			err := db.AddNotification(vLog)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
