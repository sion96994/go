package details

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ItemDetailResp
type ItemDetailResp struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

// GetItemDetailImages 获取商品详情图片
func GetItemDetailImages(apiKey, itemID string) (*ItemDetailResp, error) {
	params := url.Values{}
	params.Add("apikey", apiKey)
	params.Add("id", itemID)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest(
		"POST",
		"http://api.ds.dingdanxia.com/shop/good_images",
		body,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 请求client
	client := &http.Client{}

	// 发起请求
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//tp.Infof("resp-> %s", string(resp))

	// 解析数据
	var (
		itemResp *ItemDetailResp
	)

	err = json.Unmarshal(resp, &itemResp)
	if err != nil {
		return nil, err
	}
	if itemResp.Code != 200 {
		return nil, fmt.Errorf("获取商品详情出错, id-> %s", itemID)
	}
	return itemResp, nil
}
