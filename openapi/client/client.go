package client

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

// ClientConfig Config
type ClientConfig struct {
	// ServerURL 请求url
	ServerURL string
	// Format format: json
	Format string
	// Version api version
	Version string
	// SignMethod 加密类型(md5)
	SignMethod string
	// AppKey AppKey
	AppKey string
	// AppSecret AppSecret
	AppSecret string
}

// Client
type Client struct {
	Config ClientConfig
}

// Execute Execute
func (c *Client) Execute(params url.Values) (*simplejson.Json, error) {
	// 1. 处理参数
	// 基础参数
	params.Set("app_key", c.Config.AppKey)
	params.Set("format", c.Config.Format)
	params.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))

	// 可选参数
	params.Set("v", c.Config.Version)
	params.Set("sign_method", c.Config.SignMethod)

	// 签名
	params.Set("sign", md5Signature(c.Config.AppSecret, params))

	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", c.Config.ServerURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 加密
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// 请求client
	client := &http.Client{
		Transport: transport,
	}

	// 发起请求
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// 转换数据
	jsonData, err := simplejson.NewJson(data)
	if err != nil {
		return jsonData, err
	}
	return jsonData, nil
}

// GetClient GetClient
func GetClient(appKey, appSecret, serverURL string) *Client {
	cfg := ClientConfig{
		AppKey:     appKey,
		AppSecret:  appSecret,
		Format:     "json",
		Version:    "2.0",
		SignMethod: "md5",
		ServerURL:  serverURL,
	}
	if len(serverURL) == 0 {
		cfg.ServerURL = "http://gw.api.taobao.com/router/rest"
	}
	return &Client{
		Config: cfg,
	}
}

// md5Signature
func md5Signature(secret string, params url.Values) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := secret
	for _, k := range keys {
		v := params.Get(k)
		if k != "" && v != "" {
			str += k + v
		}
	}
	str += secret
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// IsErrorResponseExists
func IsErrorResponseExists(jsonData *simplejson.Json) error {
	if jsonData == nil {
		return errors.New("jsonData is nil")
	}
	// 错误信息
	errResp, exist := jsonData.CheckGet("error_response")
	if exist {
		bs, _ := errResp.MarshalJSON()
		return errors.New(string(bs))
	}
	return nil
}
