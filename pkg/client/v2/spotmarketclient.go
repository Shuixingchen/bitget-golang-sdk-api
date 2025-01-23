package v2

import (
	"github.com/Shuixingchen/bitget-golang-sdk-api/config"
	"github.com/Shuixingchen/bitget-golang-sdk-api/internal"
	"github.com/Shuixingchen/bitget-golang-sdk-api/internal/common"
)

type SpotMarketClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotMarketClient) Init(config config.Config) *SpotMarketClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init(config)
	return p
}

func (p *SpotMarketClient) Coins() (string, error) {
	params := internal.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/coins", params)
	return resp, err
}

func (p *SpotMarketClient) Symbols(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/public/symbols", params)
	return resp, err
}

func (p *SpotMarketClient) Fills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/fills", params)
	return resp, err
}

func (p *SpotMarketClient) Orderbook(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/orderbook", params)
	return resp, err
}

func (p *SpotMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/tickers", params)
	return resp, err
}

func (p *SpotMarketClient) Candles(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/market/candles", params)
	return resp, err
}
