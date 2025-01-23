package common

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Shuixingchen/bitget-golang-sdk-api/config"
	"github.com/Shuixingchen/bitget-golang-sdk-api/constants"
	"github.com/Shuixingchen/bitget-golang-sdk-api/internal"
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
	timesStamp := internal.TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	if constants.RSA == p.Config.SignType {
		sign = p.Signer.SignByRSA(constants.POST, uri, params, timesStamp)
	}
	requestUrl := p.Config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	internal.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
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
	timesStamp := internal.TimesStamp()
	body := internal.BuildGetParams(params)
	//fmt.Println(body)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	internal.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)

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
