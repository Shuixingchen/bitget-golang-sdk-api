package v1

import (
	"github.com/Shuixingchen/bitget-golang-sdk-api/commons/common"
	"github.com/Shuixingchen/bitget-golang-sdk-api/config"
)

type MixMarketClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *MixMarketClient) Init(config config.Config) *MixMarketClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init(config)
	return p
}

func (p *MixMarketClient) Contracts(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/mix/v1/market/contracts", params)
	return resp, err
}

func (p *MixMarketClient) Depth(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/mix/v1/market/depth", params)
	return resp, err
}

func (p *MixMarketClient) Ticker(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/mix/v1/market/ticker", params)
	return resp, err
}

func (p *MixMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/mix/v1/market/tickers", params)
	return resp, err
}

func (p *MixMarketClient) Fills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/mix/v1/market/fills", params)
	return resp, err
}

func (p *MixMarketClient) Candles(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/mix/v1/market/candles", params)
	return resp, err
}
