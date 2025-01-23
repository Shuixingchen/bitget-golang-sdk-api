package main

import (
	"fmt"

	"github.com/Shuixingchen/bitget-golang-sdk-api/config"
	"github.com/Shuixingchen/bitget-golang-sdk-api/constants"
	"github.com/Shuixingchen/bitget-golang-sdk-api/internal/model"
	"github.com/Shuixingchen/bitget-golang-sdk-api/pkg/client/ws"
)

var conf = config.Config{
	BaseUrl:       "https://api.bitget.com",
	WsUrl:         "wss://ws.bitget.com/mix/v1/stream",
	ApiKey:        "",
	SecretKey:     "",
	PASSPHRASE:    "",
	TimeoutSecond: 30,
	SignType:      constants.SHA256,
}

func main() {
	client := new(ws.BitgetWsClient).Init(true, func(message string) {
		fmt.Println("default error:" + message)
	}, func(message string) {
		fmt.Println("default error:" + message)
	}, conf)

	var channelsDef []model.SubscribeReq
	subReqDef1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channelsDef = append(channelsDef, subReqDef1)
	client.SubscribeDef(channelsDef)

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		fmt.Println("appoint:" + message)
	})
	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()
}
