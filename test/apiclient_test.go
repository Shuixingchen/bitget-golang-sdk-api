package test

import (
	"fmt"
	"testing"

	"github.com/Shuixingchen/bitget-golang-sdk-api/commons"
	"github.com/Shuixingchen/bitget-golang-sdk-api/config"
	"github.com/Shuixingchen/bitget-golang-sdk-api/constants"
	"github.com/Shuixingchen/bitget-golang-sdk-api/pkg/client"
	v1 "github.com/Shuixingchen/bitget-golang-sdk-api/pkg/client/v1"
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

func Test_PlaceOrder(t *testing.T) {
	client := new(v1.MixOrderClient).Init(conf)

	params := commons.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := client.PlaceOrder(params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_post(t *testing.T) {
	client := new(client.BitgetApiClient).Init(conf)

	params := commons.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := client.Post("/api/mix/v1/order/placeOrder", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get(t *testing.T) {

	client := new(client.BitgetApiClient).Init(conf)

	params := commons.NewParams()
	params["productType"] = "umcbl"

	resp, err := client.Get("/api/mix/v1/account/accounts", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_params(t *testing.T) {
	client := new(client.BitgetApiClient).Init(conf)

	params := commons.NewParams()

	resp, err := client.Get("/api/spot/v1/account/getInfo", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_encode_params(t *testing.T) {
	client := new(client.BitgetApiClient).Init(conf)

	params := commons.NewParams()
	params["symbol"] = "$AIUSDT"
	params["businessType"] = "spot"

	resp, err := client.Get("/api/v2/common/trade-rate", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
