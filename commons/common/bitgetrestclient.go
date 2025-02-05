package common

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Shuixingchen/bitget-golang-sdk-api/commons"
	"github.com/Shuixingchen/bitget-golang-sdk-api/config"
	"github.com/Shuixingchen/bitget-golang-sdk-api/constants"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   http.Client
	Signer       *Signer
	Config       *config.Config
}

func (p *BitgetRestClient) Init(config config.Config) *BitgetRestClient {
	p.ApiKey = config.ApiKey
	p.ApiSecretKey = config.SecretKey
	p.BaseUrl = config.BaseUrl
	p.Passphrase = config.PASSPHRASE
	p.Signer = new(Signer).Init(config.SecretKey)
	p.Config = &config
	p.HttpClient = http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}
	return p
}

func (p *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := commons.TimesStamp()
	//body, _ := commons.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	if constants.RSA == p.Config.SignType {
		sign = p.Signer.SignByRSA(constants.POST, uri, params, timesStamp)
	}
	requestUrl := p.Config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	commons.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := commons.TimesStamp()
	body := commons.BuildGetParams(params)
	//fmt.Println(body)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	commons.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
