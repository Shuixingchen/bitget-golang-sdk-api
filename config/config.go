package config

// const (
// 	BaseUrl = "https://api.bitget.com"
// 	WsUrl   = "wss://ws.bitget.com/mix/v1/stream"

// 	ApiKey        = ""
// 	SecretKey     = ""
// 	PASSPHRASE    = ""
// 	TimeoutSecond = 30
// 	SignType      = constants.SHA256
// )

type Config struct {
	ApiKey        string
	SecretKey     string
	PASSPHRASE    string
	TimeoutSecond int
	SignType      string
	BaseUrl       string
	WsUrl         string
}
